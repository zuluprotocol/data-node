// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: commands/v1/commands.proto

package v1

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/data-node/proto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *OrderSubmission) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	if !(this.Size > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Size_", fmt.Errorf(`value '%v' must be greater than '0'`, this.Size))
	}
	if this.PeggedOrder != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeggedOrder); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeggedOrder", err)
		}
	}
	return nil
}
func (this *OrderCancellation) Validate() error {
	return nil
}
func (this *OrderAmendment) Validate() error {
	if this.OrderId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrderId", fmt.Errorf(`value '%v' must not be an empty string`, this.OrderId))
	}
	if this.Price != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Price); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Price", err)
		}
	}
	if this.ExpiresAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExpiresAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExpiresAt", err)
		}
	}
	if this.PeggedOffset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PeggedOffset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PeggedOffset", err)
		}
	}
	return nil
}
func (this *LiquidityProvisionSubmission) Validate() error {
	if this.MarketId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketId", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketId))
	}
	for _, item := range this.Sells {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sells", err)
			}
		}
	}
	for _, item := range this.Buys {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Buys", err)
			}
		}
	}
	return nil
}
func (this *WithdrawSubmission) Validate() error {
	if this.Ext != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Ext); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Ext", err)
		}
	}
	return nil
}
func (this *ProposalSubmission) Validate() error {
	if nil == this.Terms {
		return github_com_mwitkow_go_proto_validators.FieldError("Terms", fmt.Errorf("message must exist"))
	}
	if this.Terms != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Terms); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Terms", err)
		}
	}
	return nil
}
func (this *VoteSubmission) Validate() error {
	if this.ProposalId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProposalId", fmt.Errorf(`value '%v' must not be an empty string`, this.ProposalId))
	}
	return nil
}
