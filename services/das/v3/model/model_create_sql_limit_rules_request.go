package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlLimitRulesRequest Request Object
type CreateSqlLimitRulesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *CreateSqlLimitRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *CreateSqlLimitRulesBody `json:"body,omitempty"`
}

func (o CreateSqlLimitRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitRulesRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitRulesRequest", string(data)}, " ")
}

type CreateSqlLimitRulesRequestXLanguage struct {
	value string
}

type CreateSqlLimitRulesRequestXLanguageEnum struct {
	ZH_CN CreateSqlLimitRulesRequestXLanguage
	EN_US CreateSqlLimitRulesRequestXLanguage
}

func GetCreateSqlLimitRulesRequestXLanguageEnum() CreateSqlLimitRulesRequestXLanguageEnum {
	return CreateSqlLimitRulesRequestXLanguageEnum{
		ZH_CN: CreateSqlLimitRulesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateSqlLimitRulesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateSqlLimitRulesRequestXLanguage) Value() string {
	return c.value
}

func (c CreateSqlLimitRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlLimitRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
