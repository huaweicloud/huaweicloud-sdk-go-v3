package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 计费类型信息，支持按需和包周期。
type OpenGaussChargeInfoListResponse struct {
	// 计费模式。  取值范围：  postPaid：后付费，即按需付费。  prePaid：预付费，即包年/包月。

	ChargeMode OpenGaussChargeInfoListResponseChargeMode `json:"charge_mode"`
}

func (o OpenGaussChargeInfoListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussChargeInfoListResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussChargeInfoListResponse", string(data)}, " ")
}

type OpenGaussChargeInfoListResponseChargeMode struct {
	value string
}

type OpenGaussChargeInfoListResponseChargeModeEnum struct {
	POST_PAID OpenGaussChargeInfoListResponseChargeMode
}

func GetOpenGaussChargeInfoListResponseChargeModeEnum() OpenGaussChargeInfoListResponseChargeModeEnum {
	return OpenGaussChargeInfoListResponseChargeModeEnum{
		POST_PAID: OpenGaussChargeInfoListResponseChargeMode{
			value: "postPaid",
		},
	}
}

func (c OpenGaussChargeInfoListResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussChargeInfoListResponseChargeMode) UnmarshalJSON(b []byte) error {
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
