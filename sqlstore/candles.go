package sqlstore

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"code.vegaprotocol.io/data-node/candlesv2"

	"github.com/shopspring/decimal"

	"code.vegaprotocol.io/data-node/entities"
	"github.com/georgysavva/scany/pgxscan"
)

const (
	sourceDataTableName    = "trades"
	candlesViewNamePrePend = sourceDataTableName + "_candle_"
)

type Candles struct {
	*SQLStore

	config candlesv2.CandleStoreConfig
	ctx    context.Context
}

func NewCandles(ctx context.Context, sqlStore *SQLStore, config candlesv2.CandleStoreConfig) (*Candles, error) {
	cs := &Candles{
		SQLStore: sqlStore,
		ctx:      ctx,
		config:   config,
	}

	for _, interval := range strings.Split(config.DefaultCandleIntervals, ",") {
		if interval == "" {
			continue
		}

		viewAlreadyExists, _, err := cs.viewExistsForInterval(ctx, interval)
		if err != nil {
			return nil, fmt.Errorf("creating candle store:%w", err)
		}

		if viewAlreadyExists {
			continue
		}

		err = cs.createViewForInterval(ctx, interval)
		if err != nil {
			return nil, fmt.Errorf("creating candles store:%w", err)
		}
	}

	return cs, nil
}

// GetCandleDataForTimeSpan gets the candles for a given interval, from and to are optional
func (cs *Candles) GetCandleDataForTimeSpan(ctx context.Context, candleId string, from *time.Time, to *time.Time,
	p entities.Pagination) ([]entities.Candle, error,
) {
	descriptor, err := candleDescriptorFromCandleId(candleId)
	if err != nil {
		return nil, fmt.Errorf("getting candle data for time span:%w", err)
	}

	exists, err := cs.CandleExists(ctx, descriptor.id)
	if err != nil {
		return nil, fmt.Errorf("getting candles for time span:%w", err)
	}

	if !exists {
		return nil, fmt.Errorf("no candle exists for candle id:%s", candleId)
	}

	var candles []entities.Candle

	query := fmt.Sprintf("SELECT period_start, open, close, high, low, volume, last_update_in_period FROM %s WHERE market_id = $1",
		descriptor.view)

	marketAsBytes, err := hex.DecodeString(descriptor.market)
	if err != nil {
		return nil, fmt.Errorf("invalid market:%w", err)
	}

	args := []interface{}{marketAsBytes}

	if from != nil {
		query = fmt.Sprintf("%s AND period_start >= %s", query, nextBindVar(&args, from))
	}

	if to != nil {
		query = fmt.Sprintf("%s AND period_start < %s", query, nextBindVar(&args, to))
	}

	query, args = orderAndPaginateQuery(query, []string{"period_start"}, p, args...)

	err = pgxscan.Select(ctx, cs.pool, &candles, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying candles: %w", err)
	}

	return candles, nil
}

// GetCandlesForMarket returns a map of existing intervals to candle ids for the given market
func (cs *Candles) GetCandlesForMarket(ctx context.Context, market string) (map[string]string, error) {
	intervalToView, err := cs.getIntervalToView(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting existing candles:%w", err)
	}

	candles := map[string]string{}
	for interval := range intervalToView {
		candles[interval] = candleDescriptorFromIntervalAndMarket(interval, market).id
	}
	return candles, nil
}

func (cs *Candles) GetCandleIdForIntervalAndMarket(ctx context.Context, interval string, market string) (bool, string, error) {
	interval, err := cs.normaliseInterval(ctx, interval)
	if err != nil {
		return false, "", fmt.Errorf("invalid interval: %w", err)
	}

	viewAlreadyExists, existingInterval, err := cs.viewExistsForInterval(ctx, interval)
	if err != nil {
		return false, "", fmt.Errorf("checking for existing view: %w", err)
	}

	if viewAlreadyExists {
		descriptor := candleDescriptorFromIntervalAndMarket(existingInterval, market)
		return true, descriptor.id, nil
	}

	return false, "", nil
}

func (cs *Candles) CreateCandleForIntervalAndMarket(ctx context.Context, interval string, market string) (string, error) {
	interval, err := cs.normaliseInterval(ctx, interval)
	if err != nil {
		return "", fmt.Errorf("invalid interval: %w", err)
	}

	viewAlreadyExists, existingInterval, err := cs.viewExistsForInterval(ctx, interval)
	if err != nil {
		return "", fmt.Errorf("checking for existing view: %w", err)
	}
	if viewAlreadyExists {
		return "",
			fmt.Errorf("an equivalent candle for interval %s already exists.  Existing equivalent candle interval:%s", interval, existingInterval)
	}

	descriptor := candleDescriptorFromIntervalAndMarket(interval, market)
	err = cs.createViewForInterval(ctx, descriptor.interval)
	if err != nil {
		return "", fmt.Errorf("creating view for interval:%w", err)
	}

	return descriptor.id, nil
}

func (cs *Candles) getIntervalToView(ctx context.Context) (map[string]string, error) {
	query := fmt.Sprintf("SELECT view_name FROM timescaledb_information.continuous_aggregates where view_name like '%s%%'",
		candlesViewNamePrePend)
	rows, err := cs.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("fetching existing views for interval: %w", err)
	}

	var viewNames []string
	for rows.Next() {
		var viewName string
		err := rows.Scan(&viewName)
		if err != nil {
			return nil, fmt.Errorf("fetching existing views for interval: %w", err)
		}
		viewNames = append(viewNames, viewName)
	}

	result := map[string]string{}
	for _, viewName := range viewNames {
		interval, err := getIntervalFromViewName(viewName)
		if err != nil {
			return nil, fmt.Errorf("fetching existing views for interval: %w", err)
		}

		result[interval] = viewName
	}
	return result, nil
}

func (cs *Candles) createViewForInterval(ctx context.Context, interval string) error {
	view := getViewNameForInterval(interval)

	query := fmt.Sprintf(`CREATE MATERIALIZED VIEW %s
		WITH (timescaledb.continuous) AS
		SELECT market_id,
			timescaledb_experimental.time_bucket_ng('%s', synthetic_time,'UTC') AS period_start, first(price, synthetic_time) AS open, last(price, synthetic_time) AS close,
			max(price) AS high, min(price) AS low, sum(size) AS volume, last(synthetic_time, synthetic_time) AS last_update_in_period
		FROM %s
		GROUP BY market_id, period_start`, view, interval, sourceDataTableName)

	_, err := cs.pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("creating view %s: %w", view, err)
	}

	return nil
}

func (cs *Candles) CandleExists(ctx context.Context, candleId string) (bool, error) {
	descriptor, err := candleDescriptorFromCandleId(candleId)
	if err != nil {
		return false, fmt.Errorf("candle exists:%w", err)
	}

	exists, _, err := cs.viewExistsForInterval(ctx, descriptor.interval)
	if err != nil {
		return false, fmt.Errorf("candle exists:%w", err)
	}

	return exists, nil
}

func (cs *Candles) viewExistsForInterval(ctx context.Context, interval string) (bool, string, error) {
	intervalToView, err := cs.getIntervalToView(ctx)
	if err != nil {
		return false, "", fmt.Errorf("checking if view exist for interval:%w", err)
	}

	if _, ok := intervalToView[interval]; ok {
		return true, interval, nil
	}

	// Also check for existing Intervals that are specified differently but amount to the same thing  (i.e 7 days = 1 week)
	existingIntervals := map[int64]string{}
	for existingInterval := range intervalToView {
		seconds, err := cs.getIntervalSeconds(ctx, existingInterval)
		if err != nil {
			return false, "", fmt.Errorf("checking if view exists for interval:%w", err)
		}
		existingIntervals[seconds] = existingInterval
	}

	seconds, err := cs.getIntervalSeconds(ctx, interval)
	if err != nil {
		return false, "", fmt.Errorf("checking if view exists for interval:%w", err)
	}

	if existingInterval, ok := existingIntervals[seconds]; ok {
		return true, existingInterval, nil
	}

	return false, "", nil
}

func (cs *Candles) normaliseInterval(ctx context.Context, interval string) (string, error) {
	var normalizedInterval string

	_, err := cs.pool.Exec(ctx, "SET intervalstyle = 'postgres_verbose' ")
	if err != nil {
		return "", fmt.Errorf("normalising interval, failed to set interval style:%w", err)
	}

	query := fmt.Sprintf("select cast( INTERVAL '%s' as text)", interval)
	row := cs.pool.QueryRow(ctx, query)

	err = row.Scan(&normalizedInterval)
	if err != nil {
		return "", fmt.Errorf("normalising interval:%s :%w", interval, err)
	}

	normalizedInterval = strings.ReplaceAll(normalizedInterval, "@ ", "")

	return normalizedInterval, nil
}

func (cs *Candles) getIntervalSeconds(ctx context.Context, interval string) (int64, error) {
	var seconds decimal.Decimal

	query := fmt.Sprintf("SELECT EXTRACT(epoch FROM INTERVAL '%s')", interval)
	row := cs.pool.QueryRow(ctx, query)

	err := row.Scan(&seconds)
	if err != nil {
		return 0, err
	}

	return seconds.IntPart(), nil
}

func getIntervalFromViewName(viewName string) (string, error) {
	split := strings.Split(viewName, candlesViewNamePrePend)
	if len(split) != 2 {
		return "", fmt.Errorf("view name has unexpected format:%s", viewName)
	}
	return strings.ReplaceAll(split[1], "_", " "), nil
}

func getViewNameForInterval(interval string) string {
	return candlesViewNamePrePend + strings.ReplaceAll(interval, " ", "_")
}

type candleDescriptor struct {
	id       string
	view     string
	interval string
	market   string
}

func candleDescriptorFromCandleId(id string) (candleDescriptor, error) {
	idx := strings.LastIndex(id, "_")

	if idx == -1 {
		return candleDescriptor{}, fmt.Errorf("invalid candle id:%s", id)
	}

	market := id[idx+1:]
	view := id[:idx]

	split := strings.Split(view, candlesViewNamePrePend)
	if len(split) != 2 {
		return candleDescriptor{}, fmt.Errorf("parsing candle id, view name has unexpected format:%s", id)
	}

	interval, err := getIntervalFromViewName(view)
	if err != nil {
		return candleDescriptor{}, fmt.Errorf("parsing candleDescriptor id, failed to get interval from view name:%w", err)
	}

	return candleDescriptor{
		id:       id,
		view:     view,
		interval: interval,
		market:   market,
	}, nil
}

func candleDescriptorFromIntervalAndMarket(interval string, market string) candleDescriptor {
	view := getViewNameForInterval(interval)
	id := view + "_" + market

	return candleDescriptor{
		id:       id,
		view:     view,
		interval: interval,
		market:   market,
	}
}