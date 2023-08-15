package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlExplainRequest Request Object
type ShowSqlExplainRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ShowSqlExplainRequestXLanguage `json:"X-Language,omitempty"`

	Body *QuerySqlPlanBody `json:"body,omitempty"`
}

func (o ShowSqlExplainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlExplainRequest struct{}"
	}

	return strings.Join([]string{"ShowSqlExplainRequest", string(data)}, " ")
}

type ShowSqlExplainRequestXLanguage struct {
	value string
}

type ShowSqlExplainRequestXLanguageEnum struct {
	ZH_CN ShowSqlExplainRequestXLanguage
	EN_US ShowSqlExplainRequestXLanguage
}

func GetShowSqlExplainRequestXLanguageEnum() ShowSqlExplainRequestXLanguageEnum {
	return ShowSqlExplainRequestXLanguageEnum{
		ZH_CN: ShowSqlExplainRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowSqlExplainRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowSqlExplainRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSqlExplainRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlExplainRequestXLanguage) UnmarshalJSON(b []byte) error {
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
