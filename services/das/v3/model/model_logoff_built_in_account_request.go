package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LogoffBuiltInAccountRequest Request Object
type LogoffBuiltInAccountRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *LogoffBuiltInAccountRequestXLanguage `json:"X-Language,omitempty"`

	Body *LogoffBuiltInAccountRequestBody `json:"body,omitempty"`
}

func (o LogoffBuiltInAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogoffBuiltInAccountRequest struct{}"
	}

	return strings.Join([]string{"LogoffBuiltInAccountRequest", string(data)}, " ")
}

type LogoffBuiltInAccountRequestXLanguage struct {
	value string
}

type LogoffBuiltInAccountRequestXLanguageEnum struct {
	ZH_CN LogoffBuiltInAccountRequestXLanguage
	EN_US LogoffBuiltInAccountRequestXLanguage
}

func GetLogoffBuiltInAccountRequestXLanguageEnum() LogoffBuiltInAccountRequestXLanguageEnum {
	return LogoffBuiltInAccountRequestXLanguageEnum{
		ZH_CN: LogoffBuiltInAccountRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: LogoffBuiltInAccountRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c LogoffBuiltInAccountRequestXLanguage) Value() string {
	return c.value
}

func (c LogoffBuiltInAccountRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LogoffBuiltInAccountRequestXLanguage) UnmarshalJSON(b []byte) error {
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
