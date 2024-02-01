package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GlobalConnectionBandwidthSizeRange struct {

	// 描述计费类型，描述可选计费类型。 取值范围：     bwd: 按带宽计费     95: 按传统型95计费
	ChargeMode *GlobalConnectionBandwidthSizeRangeChargeMode `json:"charge_mode,omitempty"`

	// 全域互联带宽最小值，单位Mbit/s。
	Min *int32 `json:"min,omitempty"`

	// 全域互联带宽最大值，单位Mbit/s。
	Max *int32 `json:"max,omitempty"`
}

func (o GlobalConnectionBandwidthSizeRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalConnectionBandwidthSizeRange struct{}"
	}

	return strings.Join([]string{"GlobalConnectionBandwidthSizeRange", string(data)}, " ")
}

type GlobalConnectionBandwidthSizeRangeChargeMode struct {
	value string
}

type GlobalConnectionBandwidthSizeRangeChargeModeEnum struct {
	BWD  GlobalConnectionBandwidthSizeRangeChargeMode
	E_95 GlobalConnectionBandwidthSizeRangeChargeMode
}

func GetGlobalConnectionBandwidthSizeRangeChargeModeEnum() GlobalConnectionBandwidthSizeRangeChargeModeEnum {
	return GlobalConnectionBandwidthSizeRangeChargeModeEnum{
		BWD: GlobalConnectionBandwidthSizeRangeChargeMode{
			value: "bwd",
		},
		E_95: GlobalConnectionBandwidthSizeRangeChargeMode{
			value: "95",
		},
	}
}

func (c GlobalConnectionBandwidthSizeRangeChargeMode) Value() string {
	return c.value
}

func (c GlobalConnectionBandwidthSizeRangeChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlobalConnectionBandwidthSizeRangeChargeMode) UnmarshalJSON(b []byte) error {
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
