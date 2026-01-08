package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTimeZoneRequest Request Object
type UpdateTimeZoneRequest struct {

	// 语言
	XLanguage *UpdateTimeZoneRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateTimeZoneRequestBody `json:"body,omitempty"`
}

func (o UpdateTimeZoneRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimeZoneRequest struct{}"
	}

	return strings.Join([]string{"UpdateTimeZoneRequest", string(data)}, " ")
}

type UpdateTimeZoneRequestXLanguage struct {
	value string
}

type UpdateTimeZoneRequestXLanguageEnum struct {
	ZH_CN UpdateTimeZoneRequestXLanguage
	EN_US UpdateTimeZoneRequestXLanguage
}

func GetUpdateTimeZoneRequestXLanguageEnum() UpdateTimeZoneRequestXLanguageEnum {
	return UpdateTimeZoneRequestXLanguageEnum{
		ZH_CN: UpdateTimeZoneRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateTimeZoneRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateTimeZoneRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateTimeZoneRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTimeZoneRequestXLanguage) UnmarshalJSON(b []byte) error {
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
