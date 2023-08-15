package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BssParamForResizeVolume 包周期扩容计费策略参数。
type BssParamForResizeVolume struct {

	// 功能说明：是否立即支付。该参数只有在云硬盘为包周期的情况下有意义。默认值为false 取值范围： * true：立即支付，从帐户余额中自动扣费 * false：不立即支付，创建订单暂不支付
	IsAutoPay *BssParamForResizeVolumeIsAutoPay `json:"isAutoPay,omitempty"`
}

func (o BssParamForResizeVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParamForResizeVolume struct{}"
	}

	return strings.Join([]string{"BssParamForResizeVolume", string(data)}, " ")
}

type BssParamForResizeVolumeIsAutoPay struct {
	value string
}

type BssParamForResizeVolumeIsAutoPayEnum struct {
	FALSE BssParamForResizeVolumeIsAutoPay
	TRUE  BssParamForResizeVolumeIsAutoPay
}

func GetBssParamForResizeVolumeIsAutoPayEnum() BssParamForResizeVolumeIsAutoPayEnum {
	return BssParamForResizeVolumeIsAutoPayEnum{
		FALSE: BssParamForResizeVolumeIsAutoPay{
			value: "false",
		},
		TRUE: BssParamForResizeVolumeIsAutoPay{
			value: "true",
		},
	}
}

func (c BssParamForResizeVolumeIsAutoPay) Value() string {
	return c.value
}

func (c BssParamForResizeVolumeIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForResizeVolumeIsAutoPay) UnmarshalJSON(b []byte) error {
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
