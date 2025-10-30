package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SelectGroupAndStreamRequest Request Object
type SelectGroupAndStreamRequest struct {

	// 请求语言类型。
	XLanguage *SelectGroupAndStreamRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o SelectGroupAndStreamRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectGroupAndStreamRequest struct{}"
	}

	return strings.Join([]string{"SelectGroupAndStreamRequest", string(data)}, " ")
}

type SelectGroupAndStreamRequestXLanguage struct {
	value string
}

type SelectGroupAndStreamRequestXLanguageEnum struct {
	EN_US SelectGroupAndStreamRequestXLanguage
	ZH_CN SelectGroupAndStreamRequestXLanguage
}

func GetSelectGroupAndStreamRequestXLanguageEnum() SelectGroupAndStreamRequestXLanguageEnum {
	return SelectGroupAndStreamRequestXLanguageEnum{
		EN_US: SelectGroupAndStreamRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: SelectGroupAndStreamRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c SelectGroupAndStreamRequestXLanguage) Value() string {
	return c.value
}

func (c SelectGroupAndStreamRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SelectGroupAndStreamRequestXLanguage) UnmarshalJSON(b []byte) error {
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
