package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type StartPipelineRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *StartPipelineRequestXLanguage `json:"X-Language,omitempty"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`
}

func (o StartPipelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPipelineRequest struct{}"
	}

	return strings.Join([]string{"StartPipelineRequest", string(data)}, " ")
}

type StartPipelineRequestXLanguage struct {
	value string
}

type StartPipelineRequestXLanguageEnum struct {
	ZH_CN StartPipelineRequestXLanguage
	EN_US StartPipelineRequestXLanguage
}

func GetStartPipelineRequestXLanguageEnum() StartPipelineRequestXLanguageEnum {
	return StartPipelineRequestXLanguageEnum{
		ZH_CN: StartPipelineRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartPipelineRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartPipelineRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartPipelineRequestXLanguage) UnmarshalJSON(b []byte) error {
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
