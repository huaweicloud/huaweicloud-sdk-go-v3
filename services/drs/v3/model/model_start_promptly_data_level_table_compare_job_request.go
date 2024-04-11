package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartPromptlyDataLevelTableCompareJobRequest Request Object
type StartPromptlyDataLevelTableCompareJobRequest struct {

	// 请求语言类型。
	XLanguage *StartPromptlyDataLevelTableCompareJobRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	// 对比任务ID。
	CompareJobId string `json:"compare_job_id"`
}

func (o StartPromptlyDataLevelTableCompareJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPromptlyDataLevelTableCompareJobRequest struct{}"
	}

	return strings.Join([]string{"StartPromptlyDataLevelTableCompareJobRequest", string(data)}, " ")
}

type StartPromptlyDataLevelTableCompareJobRequestXLanguage struct {
	value string
}

type StartPromptlyDataLevelTableCompareJobRequestXLanguageEnum struct {
	EN_US StartPromptlyDataLevelTableCompareJobRequestXLanguage
	ZH_CN StartPromptlyDataLevelTableCompareJobRequestXLanguage
}

func GetStartPromptlyDataLevelTableCompareJobRequestXLanguageEnum() StartPromptlyDataLevelTableCompareJobRequestXLanguageEnum {
	return StartPromptlyDataLevelTableCompareJobRequestXLanguageEnum{
		EN_US: StartPromptlyDataLevelTableCompareJobRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: StartPromptlyDataLevelTableCompareJobRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c StartPromptlyDataLevelTableCompareJobRequestXLanguage) Value() string {
	return c.value
}

func (c StartPromptlyDataLevelTableCompareJobRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartPromptlyDataLevelTableCompareJobRequestXLanguage) UnmarshalJSON(b []byte) error {
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
