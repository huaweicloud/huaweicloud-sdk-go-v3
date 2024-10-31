package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartMysqlCompatibilityRequest Request Object
type StartMysqlCompatibilityRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 语言。默认值：en-us。
	XLanguage *StartMysqlCompatibilityRequestXLanguage `json:"X-Language,omitempty"`

	Body *StartMySqlCompatibilityRequestBody `json:"body,omitempty"`
}

func (o StartMysqlCompatibilityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartMysqlCompatibilityRequest struct{}"
	}

	return strings.Join([]string{"StartMysqlCompatibilityRequest", string(data)}, " ")
}

type StartMysqlCompatibilityRequestXLanguage struct {
	value string
}

type StartMysqlCompatibilityRequestXLanguageEnum struct {
	ZH_CN StartMysqlCompatibilityRequestXLanguage
	EN_US StartMysqlCompatibilityRequestXLanguage
}

func GetStartMysqlCompatibilityRequestXLanguageEnum() StartMysqlCompatibilityRequestXLanguageEnum {
	return StartMysqlCompatibilityRequestXLanguageEnum{
		ZH_CN: StartMysqlCompatibilityRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartMysqlCompatibilityRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartMysqlCompatibilityRequestXLanguage) Value() string {
	return c.value
}

func (c StartMysqlCompatibilityRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartMysqlCompatibilityRequestXLanguage) UnmarshalJSON(b []byte) error {
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
