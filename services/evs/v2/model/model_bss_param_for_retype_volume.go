package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BssParamForRetypeVolume 包周期retype计费策略参数。
type BssParamForRetypeVolume struct {

	// 功能说明：是否立即支付。该参数只有在云硬盘为包周期的情况下有意义。默认值为false 取值范围： * true：立即支付，从帐户余额中自动扣费 * false：不立即支付，创建订单暂不支付
	IsAutoPay *BssParamForRetypeVolumeIsAutoPay `json:"isAutoPay,omitempty"`
}

func (o BssParamForRetypeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParamForRetypeVolume struct{}"
	}

	return strings.Join([]string{"BssParamForRetypeVolume", string(data)}, " ")
}

type BssParamForRetypeVolumeIsAutoPay struct {
	value string
}

type BssParamForRetypeVolumeIsAutoPayEnum struct {
	FALSE BssParamForRetypeVolumeIsAutoPay
	TRUE  BssParamForRetypeVolumeIsAutoPay
}

func GetBssParamForRetypeVolumeIsAutoPayEnum() BssParamForRetypeVolumeIsAutoPayEnum {
	return BssParamForRetypeVolumeIsAutoPayEnum{
		FALSE: BssParamForRetypeVolumeIsAutoPay{
			value: "false",
		},
		TRUE: BssParamForRetypeVolumeIsAutoPay{
			value: "true",
		},
	}
}

func (c BssParamForRetypeVolumeIsAutoPay) Value() string {
	return c.value
}

func (c BssParamForRetypeVolumeIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForRetypeVolumeIsAutoPay) UnmarshalJSON(b []byte) error {
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
