package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Eip Eip信息。
type Eip struct {

	// 桌面绑定的Eip的id，有值时优先绑定Eip。
	Id *string `json:"id,omitempty"`

	// EIP的类型，5_bgp（全动态BGP），5_sbgp（静态BGP）
	Type *string `json:"type,omitempty"`

	// eip带宽计费模式 - TRAFFIC：按流量计费。 - BANDWIDTH：按带宽计费。
	ChargeMode *EipChargeMode `json:"charge_mode,omitempty"`

	// 带宽大小
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`
}

func (o Eip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Eip struct{}"
	}

	return strings.Join([]string{"Eip", string(data)}, " ")
}

type EipChargeMode struct {
	value string
}

type EipChargeModeEnum struct {
	TRAFFIC   EipChargeMode
	BANDWIDTH EipChargeMode
}

func GetEipChargeModeEnum() EipChargeModeEnum {
	return EipChargeModeEnum{
		TRAFFIC: EipChargeMode{
			value: "TRAFFIC",
		},
		BANDWIDTH: EipChargeMode{
			value: "BANDWIDTH",
		},
	}
}

func (c EipChargeMode) Value() string {
	return c.value
}

func (c EipChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipChargeMode) UnmarshalJSON(b []byte) error {
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
