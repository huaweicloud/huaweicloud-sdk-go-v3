package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CheckDataFilterRulesRequest Request Object
type CheckDataFilterRulesRequest struct {

	// DRS任务的唯一标识。用于指定需要校验数据过滤规则的具体同步任务。可通过查询任务列表接口获取已有任务的job_id。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *CheckDataFilterRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *DataFilterRuleReq `json:"body,omitempty"`
}

func (o CheckDataFilterRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataFilterRulesRequest struct{}"
	}

	return strings.Join([]string{"CheckDataFilterRulesRequest", string(data)}, " ")
}

type CheckDataFilterRulesRequestXLanguage struct {
	value string
}

type CheckDataFilterRulesRequestXLanguageEnum struct {
	EN_US CheckDataFilterRulesRequestXLanguage
	ZH_CN CheckDataFilterRulesRequestXLanguage
	ES_US CheckDataFilterRulesRequestXLanguage
	PT_BR CheckDataFilterRulesRequestXLanguage
}

func GetCheckDataFilterRulesRequestXLanguageEnum() CheckDataFilterRulesRequestXLanguageEnum {
	return CheckDataFilterRulesRequestXLanguageEnum{
		EN_US: CheckDataFilterRulesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CheckDataFilterRulesRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: CheckDataFilterRulesRequestXLanguage{
			value: "es-us",
		},
		PT_BR: CheckDataFilterRulesRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c CheckDataFilterRulesRequestXLanguage) Value() string {
	return c.value
}

func (c CheckDataFilterRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckDataFilterRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
