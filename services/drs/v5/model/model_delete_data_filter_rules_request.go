package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteDataFilterRulesRequest Request Object
type DeleteDataFilterRulesRequest struct {

	// 数据迁移任务ID，用于指定要删除过滤规则的目标任务。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *DeleteDataFilterRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteDataFilterRuleReq `json:"body,omitempty"`
}

func (o DeleteDataFilterRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataFilterRulesRequest struct{}"
	}

	return strings.Join([]string{"DeleteDataFilterRulesRequest", string(data)}, " ")
}

type DeleteDataFilterRulesRequestXLanguage struct {
	value string
}

type DeleteDataFilterRulesRequestXLanguageEnum struct {
	EN_US DeleteDataFilterRulesRequestXLanguage
	ZH_CN DeleteDataFilterRulesRequestXLanguage
	ES_US DeleteDataFilterRulesRequestXLanguage
	PT_BR DeleteDataFilterRulesRequestXLanguage
}

func GetDeleteDataFilterRulesRequestXLanguageEnum() DeleteDataFilterRulesRequestXLanguageEnum {
	return DeleteDataFilterRulesRequestXLanguageEnum{
		EN_US: DeleteDataFilterRulesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteDataFilterRulesRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: DeleteDataFilterRulesRequestXLanguage{
			value: "es-us",
		},
		PT_BR: DeleteDataFilterRulesRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c DeleteDataFilterRulesRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteDataFilterRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDataFilterRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
