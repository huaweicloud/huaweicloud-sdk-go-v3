package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MarketModel struct {

	// 计费类型
	ChargeMode *MarketModelChargeMode `json:"charge_mode,omitempty"`

	PrepaidInfo *MarketModelPrepaidInfo `json:"prepaid_info,omitempty"`
}

func (o MarketModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MarketModel struct{}"
	}

	return strings.Join([]string{"MarketModel", string(data)}, " ")
}

type MarketModelChargeMode struct {
	value string
}

type MarketModelChargeModeEnum struct {
	SPOT     MarketModelChargeMode
	PREPAID  MarketModelChargeMode
	POSTPAID MarketModelChargeMode
}

func GetMarketModelChargeModeEnum() MarketModelChargeModeEnum {
	return MarketModelChargeModeEnum{
		SPOT: MarketModelChargeMode{
			value: "spot",
		},
		PREPAID: MarketModelChargeMode{
			value: "prepaid",
		},
		POSTPAID: MarketModelChargeMode{
			value: "postpaid",
		},
	}
}

func (c MarketModelChargeMode) Value() string {
	return c.value
}

func (c MarketModelChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MarketModelChargeMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
