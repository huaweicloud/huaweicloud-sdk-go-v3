package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowPipelineLastStatusV2Request struct {

	// 语言类型 中文:zh-cn 英文:en-us
	XLanguage *ShowPipelineLastStatusV2RequestXLanguage `json:"X-Language,omitempty"`

	// 流水线ID
	PipelineId string `json:"pipeline_id"`
}

func (o ShowPipelineLastStatusV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineLastStatusV2Request struct{}"
	}

	return strings.Join([]string{"ShowPipelineLastStatusV2Request", string(data)}, " ")
}

type ShowPipelineLastStatusV2RequestXLanguage struct {
	value string
}

type ShowPipelineLastStatusV2RequestXLanguageEnum struct {
	ZH_CN ShowPipelineLastStatusV2RequestXLanguage
	EN_US ShowPipelineLastStatusV2RequestXLanguage
}

func GetShowPipelineLastStatusV2RequestXLanguageEnum() ShowPipelineLastStatusV2RequestXLanguageEnum {
	return ShowPipelineLastStatusV2RequestXLanguageEnum{
		ZH_CN: ShowPipelineLastStatusV2RequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowPipelineLastStatusV2RequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowPipelineLastStatusV2RequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowPipelineLastStatusV2RequestXLanguage) UnmarshalJSON(b []byte) error {
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
