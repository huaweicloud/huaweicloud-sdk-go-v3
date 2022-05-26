package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 计费类型信息，仅支持按需。
type OpenGaussChargeInfo struct {

	// 计费模式。仅支持postPaid，后付费，即按需付费。
	ChargeMode OpenGaussChargeInfoChargeMode `json:"charge_mode"`
}

func (o OpenGaussChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussChargeInfo struct{}"
	}

	return strings.Join([]string{"OpenGaussChargeInfo", string(data)}, " ")
}

type OpenGaussChargeInfoChargeMode struct {
	value string
}

type OpenGaussChargeInfoChargeModeEnum struct {
	POST_PAID OpenGaussChargeInfoChargeMode
}

func GetOpenGaussChargeInfoChargeModeEnum() OpenGaussChargeInfoChargeModeEnum {
	return OpenGaussChargeInfoChargeModeEnum{
		POST_PAID: OpenGaussChargeInfoChargeMode{
			value: "postPaid",
		},
	}
}

func (c OpenGaussChargeInfoChargeMode) Value() string {
	return c.value
}

func (c OpenGaussChargeInfoChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussChargeInfoChargeMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
