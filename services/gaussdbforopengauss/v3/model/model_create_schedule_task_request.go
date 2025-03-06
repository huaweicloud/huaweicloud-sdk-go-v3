package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateScheduleTaskRequest Request Object
type CreateScheduleTaskRequest struct {

	// 语言。
	XLanguage *CreateScheduleTaskRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateScheduleTaskRequestBody `json:"body,omitempty"`
}

func (o CreateScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateScheduleTaskRequest", string(data)}, " ")
}

type CreateScheduleTaskRequestXLanguage struct {
	value string
}

type CreateScheduleTaskRequestXLanguageEnum struct {
	ZH_CN CreateScheduleTaskRequestXLanguage
	EN_US CreateScheduleTaskRequestXLanguage
}

func GetCreateScheduleTaskRequestXLanguageEnum() CreateScheduleTaskRequestXLanguageEnum {
	return CreateScheduleTaskRequestXLanguageEnum{
		ZH_CN: CreateScheduleTaskRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateScheduleTaskRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateScheduleTaskRequestXLanguage) Value() string {
	return c.value
}

func (c CreateScheduleTaskRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScheduleTaskRequestXLanguage) UnmarshalJSON(b []byte) error {
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
