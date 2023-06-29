package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExtendParamEip 创建弹性公网IP的extendparam字段数据结构说明
type ExtendParamEip struct {

	// 弹性公网IP的计费模式。若带宽计费类型为bandwidth，则支持prePaid和postPaid；若带宽计费类型为traffic，仅支持postPaid。取值范围：prePaid：预付费，即包年包月postPaid：后付费，即按需付费 说明：如果bandwidth对象中sharetype是WHOLE且id有值，弹性公网IP只能创建为按需付费的，故该参数传参“prePaid”无效。
	ChargingMode ExtendParamEipChargingMode `json:"chargingMode"`
}

func (o ExtendParamEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendParamEip struct{}"
	}

	return strings.Join([]string{"ExtendParamEip", string(data)}, " ")
}

type ExtendParamEipChargingMode struct {
	value string
}

type ExtendParamEipChargingModeEnum struct {
	PRE_PAID  ExtendParamEipChargingMode
	POST_PAID ExtendParamEipChargingMode
}

func GetExtendParamEipChargingModeEnum() ExtendParamEipChargingModeEnum {
	return ExtendParamEipChargingModeEnum{
		PRE_PAID: ExtendParamEipChargingMode{
			value: "prePaid",
		},
		POST_PAID: ExtendParamEipChargingMode{
			value: "postPaid",
		},
	}
}

func (c ExtendParamEipChargingMode) Value() string {
	return c.value
}

func (c ExtendParamEipChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtendParamEipChargingMode) UnmarshalJSON(b []byte) error {
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
