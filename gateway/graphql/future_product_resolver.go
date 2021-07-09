package gql

import (
	"context"

	types "code.vegaprotocol.io/data-node/proto"
)

type myFutureProductResolver VegaResolverRoot

func (r *myFutureProductResolver) SettlementAsset(ctx context.Context, obj *types.FutureProduct) (*types.Asset, error) {
	return r.r.getAssetByID(ctx, obj.SettlementAsset)
}
