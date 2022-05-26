package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 计费类型信息，支持按需，默认为按需。
type OpenGaussChargeInfoResponse struct {

	// 计费模式。 取值范围： postPaid：后付费，即按需付费。
	ChargeMode OpenGaussChargeInfoResponseChargeMode `json:"charge_mode"`
}

func (o OpenGaussChargeInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussChargeInfoResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussChargeInfoResponse", string(data)}, " ")
}

type OpenGaussChargeInfoResponseChargeMode struct {
	value string
}

type OpenGaussChargeInfoResponseChargeModeEnum struct {
	POST_PAID OpenGaussChargeInfoResponseChargeMode
}

func GetOpenGaussChargeInfoResponseChargeModeEnum() OpenGaussChargeInfoResponseChargeModeEnum {
	return OpenGaussChargeInfoResponseChargeModeEnum{
		POST_PAID: OpenGaussChargeInfoResponseChargeMode{
			value: "postPaid",
		},
	}
}

func (c OpenGaussChargeInfoResponseChargeMode) Value() string {
	return c.value
}

func (c OpenGaussChargeInfoResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussChargeInfoResponseChargeMode) UnmarshalJSON(b []byte) error {
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
