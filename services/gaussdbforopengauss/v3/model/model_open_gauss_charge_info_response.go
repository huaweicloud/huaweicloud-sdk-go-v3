package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 付费方式信息，仅支持按需。
type OpenGaussChargeInfoResponse struct {
	// 创建类型，支持按需。

	ChargeMode OpenGaussChargeInfoResponseChargeMode `json:"charge_mode"`
}

func (o OpenGaussChargeInfoResponse) String() string {
	data, err := json.Marshal(o)
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

func (c OpenGaussChargeInfoResponseChargeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
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
