package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDataFilterRulesResultRequest Request Object
type ShowDataFilterRulesResultRequest struct {

	// 数据迁移任务ID，用于指定要查询过滤规则的目标任务。
	JobId string `json:"job_id"`

	// 查询数据过滤规则结果请求ID，用于关联之前提交的查询过滤校验请求。
	QueryId *string `json:"query_id,omitempty"`

	// 请求语言类型。
	XLanguage *ShowDataFilterRulesResultRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowDataFilterRulesResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataFilterRulesResultRequest struct{}"
	}

	return strings.Join([]string{"ShowDataFilterRulesResultRequest", string(data)}, " ")
}

type ShowDataFilterRulesResultRequestXLanguage struct {
	value string
}

type ShowDataFilterRulesResultRequestXLanguageEnum struct {
	EN_US ShowDataFilterRulesResultRequestXLanguage
	ZH_CN ShowDataFilterRulesResultRequestXLanguage
	ES_US ShowDataFilterRulesResultRequestXLanguage
	PT_BR ShowDataFilterRulesResultRequestXLanguage
}

func GetShowDataFilterRulesResultRequestXLanguageEnum() ShowDataFilterRulesResultRequestXLanguageEnum {
	return ShowDataFilterRulesResultRequestXLanguageEnum{
		EN_US: ShowDataFilterRulesResultRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowDataFilterRulesResultRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: ShowDataFilterRulesResultRequestXLanguage{
			value: "es-us",
		},
		PT_BR: ShowDataFilterRulesResultRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c ShowDataFilterRulesResultRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDataFilterRulesResultRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDataFilterRulesResultRequestXLanguage) UnmarshalJSON(b []byte) error {
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
