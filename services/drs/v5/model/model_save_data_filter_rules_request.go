package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SaveDataFilterRulesRequest Request Object
type SaveDataFilterRulesRequest struct {

	// 数据迁移任务ID，用于指定要保存过滤规则的目标任务。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *SaveDataFilterRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *DataFilterRuleReq `json:"body,omitempty"`
}

func (o SaveDataFilterRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveDataFilterRulesRequest struct{}"
	}

	return strings.Join([]string{"SaveDataFilterRulesRequest", string(data)}, " ")
}

type SaveDataFilterRulesRequestXLanguage struct {
	value string
}

type SaveDataFilterRulesRequestXLanguageEnum struct {
	EN_US SaveDataFilterRulesRequestXLanguage
	ZH_CN SaveDataFilterRulesRequestXLanguage
	ES_US SaveDataFilterRulesRequestXLanguage
	PT_BR SaveDataFilterRulesRequestXLanguage
}

func GetSaveDataFilterRulesRequestXLanguageEnum() SaveDataFilterRulesRequestXLanguageEnum {
	return SaveDataFilterRulesRequestXLanguageEnum{
		EN_US: SaveDataFilterRulesRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: SaveDataFilterRulesRequestXLanguage{
			value: "zh-cn",
		},
		ES_US: SaveDataFilterRulesRequestXLanguage{
			value: "es-us",
		},
		PT_BR: SaveDataFilterRulesRequestXLanguage{
			value: "pt-br",
		},
	}
}

func (c SaveDataFilterRulesRequestXLanguage) Value() string {
	return c.value
}

func (c SaveDataFilterRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SaveDataFilterRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
