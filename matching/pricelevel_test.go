package matching

import (
	"testing"

	"code.vegaprotocol.io/data-node/logging"
	"code.vegaprotocol.io/data-node/types"
	"code.vegaprotocol.io/data-node/types/num"

	"github.com/stretchr/testify/assert"
)

func TestGetPriceLevel(t *testing.T) {
	side := &OrderBookSide{side: types.Side_SIDE_SELL}
	assert.Equal(t, 0, len(side.levels))
	side.getPriceLevel(num.NewUint(100))
	assert.Equal(t, 1, len(side.levels))

	side.getPriceLevel(num.NewUint(110))
	assert.Equal(t, 2, len(side.levels))

	side.getPriceLevel(num.NewUint(100))
	assert.Equal(t, 2, len(side.levels))
}

func TestAddAndRemoveOrdersToPriceLevel(t *testing.T) {
	side := &OrderBookSide{side: types.Side_SIDE_SELL}
	l := side.getPriceLevel(num.NewUint(100))
	order := &types.Order{
		MarketId:    "testOrderBook",
		PartyId:     "A",
		Side:        types.Side_SIDE_SELL,
		Price:       num.NewUint(101),
		Size:        100,
		Remaining:   100,
		TimeInForce: types.Order_TIME_IN_FORCE_GTC,
		CreatedAt:   0,
	}

	// add orders
	assert.Equal(t, 0, len(l.orders))
	l.addOrder(order)
	assert.Equal(t, 1, len(l.orders))
	l.addOrder(order)
	assert.Equal(t, 2, len(l.orders))

	// remove orders
	l.removeOrder(1)
	assert.Equal(t, 1, len(l.orders))
	l.removeOrder(0)
	assert.Equal(t, 0, len(l.orders))
}

func TestUncross(t *testing.T) {
	logger := logging.NewTestLogger()
	defer logger.Sync()

	side := &OrderBookSide{side: types.Side_SIDE_SELL}
	l := side.getPriceLevel(num.NewUint(100))
	passiveOrder := &types.Order{
		MarketId:    "testOrderBook",
		PartyId:     "A",
		Side:        types.Side_SIDE_SELL,
		Price:       num.NewUint(101),
		Size:        100,
		Remaining:   100,
		TimeInForce: types.Order_TIME_IN_FORCE_GTC,
		CreatedAt:   0,
	}
	l.addOrder(passiveOrder)

	aggresiveOrder := &types.Order{
		MarketId:    "testOrderBook",
		PartyId:     "B",
		Side:        types.Side_SIDE_BUY,
		Price:       num.NewUint(101),
		Size:        100,
		Remaining:   100,
		TimeInForce: types.Order_TIME_IN_FORCE_GTC,
		CreatedAt:   0,
	}
	filled, trades, impactedOrders, err := l.uncross(aggresiveOrder, true)
	assert.Equal(t, true, filled)
	assert.Equal(t, 1, len(trades))
	assert.Equal(t, 1, len(impactedOrders))
	assert.NoError(t, err)
}
