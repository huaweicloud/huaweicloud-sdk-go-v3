package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSqlLimitRulesRequest Request Object
type DeleteSqlLimitRulesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *DeleteSqlLimitRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *DeleteSqlLimitRulesBody `json:"body,omitempty"`
}

func (o DeleteSqlLimitRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitRulesRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitRulesRequest", string(data)}, " ")
}

type DeleteSqlLimitRulesRequestXLanguage struct {
	value string
}

type DeleteSqlLimitRulesRequestXLanguageEnum struct {
	ZH_CN DeleteSqlLimitRulesRequestXLanguage
	EN_US DeleteSqlLimitRulesRequestXLanguage
}

func GetDeleteSqlLimitRulesRequestXLanguageEnum() DeleteSqlLimitRulesRequestXLanguageEnum {
	return DeleteSqlLimitRulesRequestXLanguageEnum{
		ZH_CN: DeleteSqlLimitRulesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteSqlLimitRulesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteSqlLimitRulesRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteSqlLimitRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSqlLimitRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
