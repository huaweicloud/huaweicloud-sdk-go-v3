package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type UpdateReadWeightRequest struct {
	// 语言

	XLanguage *UpdateReadWeightRequestXLanguage `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ModifyProxyWeightRequest `json:"body,omitempty"`
}

func (o UpdateReadWeightRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReadWeightRequest struct{}"
	}

	return strings.Join([]string{"UpdateReadWeightRequest", string(data)}, " ")
}

type UpdateReadWeightRequestXLanguage struct {
	value string
}

type UpdateReadWeightRequestXLanguageEnum struct {
	ZH_CN UpdateReadWeightRequestXLanguage
	EN_US UpdateReadWeightRequestXLanguage
}

func GetUpdateReadWeightRequestXLanguageEnum() UpdateReadWeightRequestXLanguageEnum {
	return UpdateReadWeightRequestXLanguageEnum{
		ZH_CN: UpdateReadWeightRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateReadWeightRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateReadWeightRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateReadWeightRequestXLanguage) UnmarshalJSON(b []byte) error {
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
