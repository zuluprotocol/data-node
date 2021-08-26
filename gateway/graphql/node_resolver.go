package gql

import (
	"context"
	"fmt"

	proto "code.vegaprotocol.io/protos/vega"
)

type nodeResolver VegaResolverRoot

func (r *nodeResolver) Status(ctx context.Context, obj *proto.Node) (NodeStatus, error) {
	return nodeStatusFromProto(obj.Status)
}

func (r *nodeResolver) Delegations(ctx context.Context, obj *proto.Node, partyID *string) ([]*proto.Delegation, error) {
	if partyID == nil || *partyID == "" {
		return obj.Delagations, nil
	}

	partyDelegations := []*proto.Delegation{}

	for _, d := range obj.Delagations {
		if d.Party == *partyID {
			partyDelegations = append(partyDelegations, d)
		}
	}

	return partyDelegations, nil
}

func nodeStatusFromProto(s proto.NodeStatus) (NodeStatus, error) {
	switch s {
	case proto.NodeStatus_NODE_STATUS_VALIDATOR:
		return NodeStatusValidator, nil
	case proto.NodeStatus_NODE_STATUS_NON_VALIDATOR:
		return NodeStatusNonValidator, nil
	default:
		return NodeStatus(""), fmt.Errorf("failed to convert NodeStatus from Proto to GraphQL: %s", s.String())
	}
}
