package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSqlLimitRulesRequest Request Object
type UpdateSqlLimitRulesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *UpdateSqlLimitRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateSqlLimitRulesBody `json:"body,omitempty"`
}

func (o UpdateSqlLimitRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitRulesRequest struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitRulesRequest", string(data)}, " ")
}

type UpdateSqlLimitRulesRequestXLanguage struct {
	value string
}

type UpdateSqlLimitRulesRequestXLanguageEnum struct {
	ZH_CN UpdateSqlLimitRulesRequestXLanguage
	EN_US UpdateSqlLimitRulesRequestXLanguage
}

func GetUpdateSqlLimitRulesRequestXLanguageEnum() UpdateSqlLimitRulesRequestXLanguageEnum {
	return UpdateSqlLimitRulesRequestXLanguageEnum{
		ZH_CN: UpdateSqlLimitRulesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateSqlLimitRulesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateSqlLimitRulesRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateSqlLimitRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlLimitRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
