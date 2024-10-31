package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateMysqlCompatibilityRequest Request Object
type UpdateMysqlCompatibilityRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言。默认值：en-us。
	XLanguage *UpdateMysqlCompatibilityRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateMySqlCompatibilityRequestBody `json:"body,omitempty"`
}

func (o UpdateMysqlCompatibilityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMysqlCompatibilityRequest struct{}"
	}

	return strings.Join([]string{"UpdateMysqlCompatibilityRequest", string(data)}, " ")
}

type UpdateMysqlCompatibilityRequestXLanguage struct {
	value string
}

type UpdateMysqlCompatibilityRequestXLanguageEnum struct {
	ZH_CN UpdateMysqlCompatibilityRequestXLanguage
	EN_US UpdateMysqlCompatibilityRequestXLanguage
}

func GetUpdateMysqlCompatibilityRequestXLanguageEnum() UpdateMysqlCompatibilityRequestXLanguageEnum {
	return UpdateMysqlCompatibilityRequestXLanguageEnum{
		ZH_CN: UpdateMysqlCompatibilityRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateMysqlCompatibilityRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateMysqlCompatibilityRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateMysqlCompatibilityRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateMysqlCompatibilityRequestXLanguage) UnmarshalJSON(b []byte) error {
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
