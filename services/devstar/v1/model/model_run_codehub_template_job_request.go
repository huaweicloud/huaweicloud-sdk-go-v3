package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type RunCodehubTemplateJobRequest struct {
	// 语言类型，缺省值为“zh-cn”。  枚举值： - zh-cn：中文 - en-us：英文

	XLanguage *RunCodehubTemplateJobRequestXLanguage `json:"X-Language,omitempty"`

	Body *CodehubJobInfo `json:"body,omitempty"`
}

func (o RunCodehubTemplateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCodehubTemplateJobRequest struct{}"
	}

	return strings.Join([]string{"RunCodehubTemplateJobRequest", string(data)}, " ")
}

type RunCodehubTemplateJobRequestXLanguage struct {
	value string
}

type RunCodehubTemplateJobRequestXLanguageEnum struct {
	ZH_CN RunCodehubTemplateJobRequestXLanguage
	EN_US RunCodehubTemplateJobRequestXLanguage
}

func GetRunCodehubTemplateJobRequestXLanguageEnum() RunCodehubTemplateJobRequestXLanguageEnum {
	return RunCodehubTemplateJobRequestXLanguageEnum{
		ZH_CN: RunCodehubTemplateJobRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: RunCodehubTemplateJobRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c RunCodehubTemplateJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunCodehubTemplateJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
