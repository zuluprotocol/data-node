// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: markets.proto

package proto

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/data-node/proto/oracles/v1"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AuctionDuration) Validate() error {
	return nil
}
func (this *ContinuousTrading) Validate() error {
	return nil
}
func (this *DiscreteTrading) Validate() error {
	if !(this.DurationNs > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("DurationNs", fmt.Errorf(`value '%v' must be greater than '0'`, this.DurationNs))
	}
	if !(this.DurationNs < 2592000000000000) {
		return github_com_mwitkow_go_proto_validators.FieldError("DurationNs", fmt.Errorf(`value '%v' must be less than '2592000000000000'`, this.DurationNs))
	}
	return nil
}
func (this *Future) Validate() error {
	if this.OracleSpec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OracleSpec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OracleSpec", err)
		}
	}
	if this.OracleSpecBinding != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OracleSpecBinding); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OracleSpecBinding", err)
		}
	}
	return nil
}
func (this *OracleSpecToFutureBinding) Validate() error {
	return nil
}
func (this *InstrumentMetadata) Validate() error {
	return nil
}
func (this *Instrument) Validate() error {
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	if oneOfNester, ok := this.GetProduct().(*Instrument_Future); ok {
		if oneOfNester.Future != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Future); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Future", err)
			}
		}
	}
	return nil
}
func (this *LogNormalRiskModel) Validate() error {
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *LogNormalModelParams) Validate() error {
	return nil
}
func (this *SimpleRiskModel) Validate() error {
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *SimpleModelParams) Validate() error {
	if !(this.MaxMoveUp >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MaxMoveUp", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.MaxMoveUp))
	}
	if !(this.MinMoveDown <= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("MinMoveDown", fmt.Errorf(`value '%v' must be lower than or equal to '0'`, this.MinMoveDown))
	}
	if !(this.ProbabilityOfTrading >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProbabilityOfTrading", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.ProbabilityOfTrading))
	}
	if !(this.ProbabilityOfTrading <= 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProbabilityOfTrading", fmt.Errorf(`value '%v' must be lower than or equal to '1'`, this.ProbabilityOfTrading))
	}
	return nil
}
func (this *ScalingFactors) Validate() error {
	return nil
}
func (this *MarginCalculator) Validate() error {
	if this.ScalingFactors != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ScalingFactors); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ScalingFactors", err)
		}
	}
	return nil
}
func (this *TradableInstrument) Validate() error {
	if this.Instrument != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Instrument); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Instrument", err)
		}
	}
	if this.MarginCalculator != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarginCalculator); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarginCalculator", err)
		}
	}
	if oneOfNester, ok := this.GetRiskModel().(*TradableInstrument_LogNormalRiskModel); ok {
		if oneOfNester.LogNormalRiskModel != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LogNormalRiskModel); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LogNormalRiskModel", err)
			}
		}
	}
	if oneOfNester, ok := this.GetRiskModel().(*TradableInstrument_SimpleRiskModel); ok {
		if oneOfNester.SimpleRiskModel != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SimpleRiskModel); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SimpleRiskModel", err)
			}
		}
	}
	return nil
}
func (this *FeeFactors) Validate() error {
	return nil
}
func (this *Fees) Validate() error {
	if this.Factors != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Factors); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Factors", err)
		}
	}
	return nil
}
func (this *PriceMonitoringTrigger) Validate() error {
	if !(this.Horizon > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Horizon", fmt.Errorf(`value '%v' must be greater than '0'`, this.Horizon))
	}
	if !(this.Probability > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Probability", fmt.Errorf(`value '%v' must be strictly greater than '0'`, this.Probability))
	}
	if !(this.Probability < 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Probability", fmt.Errorf(`value '%v' must be strictly lower than '1'`, this.Probability))
	}
	if !(this.AuctionExtension > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("AuctionExtension", fmt.Errorf(`value '%v' must be greater than '0'`, this.AuctionExtension))
	}
	return nil
}
func (this *PriceMonitoringParameters) Validate() error {
	for _, item := range this.Triggers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Triggers", err)
			}
		}
	}
	return nil
}
func (this *PriceMonitoringSettings) Validate() error {
	if this.Parameters != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Parameters); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Parameters", err)
		}
	}
	return nil
}
func (this *LiquidityMonitoringParameters) Validate() error {
	if this.TargetStakeParameters != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TargetStakeParameters); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TargetStakeParameters", err)
		}
	}
	if !(this.TriggeringRatio >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TriggeringRatio", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.TriggeringRatio))
	}
	if !(this.TriggeringRatio <= 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("TriggeringRatio", fmt.Errorf(`value '%v' must be lower than or equal to '1'`, this.TriggeringRatio))
	}
	return nil
}
func (this *TargetStakeParameters) Validate() error {
	if !(this.TimeWindow > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("TimeWindow", fmt.Errorf(`value '%v' must be greater than '0'`, this.TimeWindow))
	}
	if !(this.ScalingFactor > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ScalingFactor", fmt.Errorf(`value '%v' must be strictly greater than '0'`, this.ScalingFactor))
	}
	return nil
}
func (this *Market) Validate() error {
	if this.TradableInstrument != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TradableInstrument); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TradableInstrument", err)
		}
	}
	if this.Fees != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fees); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fees", err)
		}
	}
	if this.OpeningAuction != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OpeningAuction); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OpeningAuction", err)
		}
	}
	if oneOfNester, ok := this.GetTradingModeConfig().(*Market_Continuous); ok {
		if oneOfNester.Continuous != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Continuous); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Continuous", err)
			}
		}
	}
	if oneOfNester, ok := this.GetTradingModeConfig().(*Market_Discrete); ok {
		if oneOfNester.Discrete != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Discrete); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Discrete", err)
			}
		}
	}
	if this.PriceMonitoringSettings != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PriceMonitoringSettings); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PriceMonitoringSettings", err)
		}
	}
	if this.LiquidityMonitoringParameters != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LiquidityMonitoringParameters); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LiquidityMonitoringParameters", err)
		}
	}
	if this.MarketTimestamps != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarketTimestamps); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarketTimestamps", err)
		}
	}
	return nil
}
func (this *MarketTimestamps) Validate() error {
	return nil
}
