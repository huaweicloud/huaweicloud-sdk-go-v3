package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StopJobActionRequest Request Object
type StopJobActionRequest struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *StopJobActionRequestXLanguage `json:"X-Language,omitempty"`

	Body *StopJobActionReq `json:"body,omitempty"`
}

func (o StopJobActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobActionRequest struct{}"
	}

	return strings.Join([]string{"StopJobActionRequest", string(data)}, " ")
}

type StopJobActionRequestXLanguage struct {
	value string
}

type StopJobActionRequestXLanguageEnum struct {
	EN_US StopJobActionRequestXLanguage
	ZH_CN StopJobActionRequestXLanguage
}

func GetStopJobActionRequestXLanguageEnum() StopJobActionRequestXLanguageEnum {
	return StopJobActionRequestXLanguageEnum{
		EN_US: StopJobActionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: StopJobActionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c StopJobActionRequestXLanguage) Value() string {
	return c.value
}

func (c StopJobActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopJobActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
