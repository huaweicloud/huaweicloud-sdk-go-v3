package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddFullSqlTaskRequest Request Object
type AddFullSqlTaskRequest struct {

	// 语言
	XLanguage *AddFullSqlTaskRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *AddFullSqlTaskBody `json:"body,omitempty"`
}

func (o AddFullSqlTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFullSqlTaskRequest struct{}"
	}

	return strings.Join([]string{"AddFullSqlTaskRequest", string(data)}, " ")
}

type AddFullSqlTaskRequestXLanguage struct {
	value string
}

type AddFullSqlTaskRequestXLanguageEnum struct {
	ZH_CN AddFullSqlTaskRequestXLanguage
	EN_US AddFullSqlTaskRequestXLanguage
}

func GetAddFullSqlTaskRequestXLanguageEnum() AddFullSqlTaskRequestXLanguageEnum {
	return AddFullSqlTaskRequestXLanguageEnum{
		ZH_CN: AddFullSqlTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: AddFullSqlTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c AddFullSqlTaskRequestXLanguage) Value() string {
	return c.value
}

func (c AddFullSqlTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddFullSqlTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
