package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDataProcessingRulesResultRequest Request Object
type ShowDataProcessingRulesResultRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowDataProcessingRulesResultRequestXLanguage `json:"X-Language,omitempty"`

	// 更新数据加工规则ID。
	QueryId string `json:"query_id"`
}

func (o ShowDataProcessingRulesResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataProcessingRulesResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDataProcessingRulesResultRequest", string(data)}, " ")
}

type ShowDataProcessingRulesResultRequestXLanguage struct {
	value string
}

type ShowDataProcessingRulesResultRequestXLanguageEnum struct {
	EN_US ShowDataProcessingRulesResultRequestXLanguage
	ZH_CN ShowDataProcessingRulesResultRequestXLanguage
}

func GetShowDataProcessingRulesResultRequestXLanguageEnum() ShowDataProcessingRulesResultRequestXLanguageEnum {
	return ShowDataProcessingRulesResultRequestXLanguageEnum{
		EN_US: ShowDataProcessingRulesResultRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDataProcessingRulesResultRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowDataProcessingRulesResultRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDataProcessingRulesResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataProcessingRulesResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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
