package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateJobRequest struct {

	// 请求语言类型。
	XLanguage *CreateJobRequestXLanguage `json:"X-Language,omitempty"`

	Body *SingleCreateJobReq `json:"body,omitempty"`
}

func (o CreateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobRequest struct{}"
	}

	return strings.Join([]string{"CreateJobRequest", string(data)}, " ")
}

type CreateJobRequestXLanguage struct {
	value string
}

type CreateJobRequestXLanguageEnum struct {
	EN_US CreateJobRequestXLanguage
	ZH_CN CreateJobRequestXLanguage
}

func GetCreateJobRequestXLanguageEnum() CreateJobRequestXLanguageEnum {
	return CreateJobRequestXLanguageEnum{
		EN_US: CreateJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CreateJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c CreateJobRequestXLanguage) Value() string {
	return c.value
}

func (c CreateJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
