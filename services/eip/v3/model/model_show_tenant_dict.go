package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowTenantDict struct {

	// - 功能说明：弹性公网IP的唯一标识
	Id *string `json:"id,omitempty"`

	// 带宽的计费模式
	ChargeMode *ShowTenantDictChargeMode `json:"charge_mode,omitempty"`

	// 该类型带宽可购买的最小size
	MinSize *int32 `json:"min_size,omitempty"`

	// 该类型带宽可购买的最大size
	MaxSize *int32 `json:"max_size,omitempty"`

	ExtLimit *ExtLimitPojo `json:"ext_limit,omitempty"`
}

func (o ShowTenantDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantDict struct{}"
	}

	return strings.Join([]string{"ShowTenantDict", string(data)}, " ")
}

type ShowTenantDictChargeMode struct {
	value string
}

type ShowTenantDictChargeModeEnum struct {
	BANDWIDTH ShowTenantDictChargeMode
	TRAFFIC   ShowTenantDictChargeMode
}

func GetShowTenantDictChargeModeEnum() ShowTenantDictChargeModeEnum {
	return ShowTenantDictChargeModeEnum{
		BANDWIDTH: ShowTenantDictChargeMode{
			value: "bandwidth",
		},
		TRAFFIC: ShowTenantDictChargeMode{
			value: "traffic",
		},
	}
}

func (c ShowTenantDictChargeMode) Value() string {
	return c.value
}

func (c ShowTenantDictChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTenantDictChargeMode) UnmarshalJSON(b []byte) error {
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
