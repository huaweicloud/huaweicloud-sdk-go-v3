package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectDataFilterRulesRequest Request Object
type CollectDataFilterRulesRequest struct {

	// 数据迁移任务ID，用于指定要查询过滤规则的目标任务。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *CollectDataFilterRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *QueryDataFilterRuleReq `json:"body,omitempty"`
}

func (o CollectDataFilterRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectDataFilterRulesRequest struct{}"
	}

	return strings.Join([]string{"CollectDataFilterRulesRequest", string(data)}, " ")
}

type CollectDataFilterRulesRequestXLanguage struct {
	value string
}

type CollectDataFilterRulesRequestXLanguageEnum struct {
	EN_US CollectDataFilterRulesRequestXLanguage
	ZH_CN CollectDataFilterRulesRequestXLanguage
	ES_US CollectDataFilterRulesRequestXLanguage
	PT_BR CollectDataFilterRulesRequestXLanguage
}

func GetCollectDataFilterRulesRequestXLanguageEnum() CollectDataFilterRulesRequestXLanguageEnum {
	return CollectDataFilterRulesRequestXLanguageEnum{
		EN_US: CollectDataFilterRulesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: CollectDataFilterRulesRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: CollectDataFilterRulesRequestXLanguage{
			value: "es-us",
		},
		PT_BR: CollectDataFilterRulesRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c CollectDataFilterRulesRequestXLanguage) Value() string {
	return c.value
}

func (c CollectDataFilterRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectDataFilterRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
