package sqlstore

import (
	"context"
	"fmt"
	"strings"

	"code.vegaprotocol.io/data-node/entities"
	"github.com/georgysavva/scany/pgxscan"
)

type Votes struct {
	*SQLStore
}

func NewVotes(sqlStore *SQLStore) *Votes {
	d := &Votes{
		SQLStore: sqlStore,
	}
	return d
}

func (vs *Votes) Add(ctx context.Context, v entities.Vote) error {
	_, err := vs.pool.Exec(ctx,
		`INSERT INTO votes(
			proposal_id,
			party_id,
			value,
			vega_time,
			initial_time,
			total_governance_token_balance,
			total_governance_token_weight,
			total_equity_like_share_weight
		)
		 VALUES ($1,  $2,  $3,  $4,  $5, $6, $7, $8)
		 ON CONFLICT (proposal_id, party_id, vega_time) DO UPDATE SET
			value = EXCLUDED.value,
			total_governance_token_balance =EXCLUDED.total_governance_token_balance,
			total_governance_token_weight = EXCLUDED.total_governance_token_weight,
			total_equity_like_share_weight = EXCLUDED.total_equity_like_share_weight;
		`,
		v.ProposalID, v.PartyID, v.Value, v.VegaTime, v.InitialTime,
		v.TotalGovernanceTokenBalance, v.TotalGovernanceTokenWeight, v.TotalEquityLikeShareWeight)
	return err
}

func (rs *Votes) GetYesVotesForProposal(ctx context.Context, proposalIDStr string) ([]entities.Vote, error) {
	yes := entities.VoteValueYes
	return rs.Get(ctx, &proposalIDStr, nil, &yes)
}

func (rs *Votes) GetNoVotesForProposal(ctx context.Context, proposalIDStr string) ([]entities.Vote, error) {
	no := entities.VoteValueNo
	return rs.Get(ctx, &proposalIDStr, nil, &no)
}

func (rs *Votes) GetByParty(ctx context.Context, partyIDStr string) ([]entities.Vote, error) {
	return rs.Get(ctx, nil, &partyIDStr, nil)
}

func (rs *Votes) Get(ctx context.Context,
	proposalIDStr *string,
	partyIDStr *string,
	value *entities.VoteValue,
) ([]entities.Vote, error) {
	query := `SELECT * FROM votes_current`
	args := []interface{}{}

	conditions := []string{}

	if proposalIDStr != nil {
		proposalID := entities.NewProposalID(*proposalIDStr)
		conditions = append(conditions, fmt.Sprintf("proposal_id=%s", nextBindVar(&args, proposalID)))
	}

	if partyIDStr != nil {
		partyID := entities.NewPartyID(*partyIDStr)
		conditions = append(conditions, fmt.Sprintf("party_id=%s", nextBindVar(&args, partyID)))
	}

	if value != nil {
		conditions = append(conditions, fmt.Sprintf("value=%s", nextBindVar(&args, *value)))
	}

	if len(conditions) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditions, " AND "))
	}

	votes := []entities.Vote{}
	err := pgxscan.Select(ctx, rs.pool, &votes, query, args...)
	if err != nil {
		return nil, fmt.Errorf("querying votes: %w", err)
	}
	return votes, nil
}