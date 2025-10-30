package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDasRecommendSqlLimitRuleRequest Request Object
type ShowDasRecommendSqlLimitRuleRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ShowDasRecommendSqlLimitRuleRequestXLanguage `json:"X-Language,omitempty"`

	Body *ShowRecommendSqlLimitRuleRequestBody `json:"body,omitempty"`
}

func (o ShowDasRecommendSqlLimitRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDasRecommendSqlLimitRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowDasRecommendSqlLimitRuleRequest", string(data)}, " ")
}

type ShowDasRecommendSqlLimitRuleRequestXLanguage struct {
	value string
}

type ShowDasRecommendSqlLimitRuleRequestXLanguageEnum struct {
	ZH_CN ShowDasRecommendSqlLimitRuleRequestXLanguage
	EN_US ShowDasRecommendSqlLimitRuleRequestXLanguage
}

func GetShowDasRecommendSqlLimitRuleRequestXLanguageEnum() ShowDasRecommendSqlLimitRuleRequestXLanguageEnum {
	return ShowDasRecommendSqlLimitRuleRequestXLanguageEnum{
		ZH_CN: ShowDasRecommendSqlLimitRuleRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowDasRecommendSqlLimitRuleRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowDasRecommendSqlLimitRuleRequestXLanguage) Value() string {
	return c.value
}

func (c ShowDasRecommendSqlLimitRuleRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDasRecommendSqlLimitRuleRequestXLanguage) UnmarshalJSON(b []byte) error {
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
