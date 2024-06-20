package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ParseSqlLimitRulesRequest Request Object
type ParseSqlLimitRulesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *ParseSqlLimitRulesRequestXLanguage `json:"X-Language,omitempty"`

	Body *ParseSqlLimitRulesReq `json:"body,omitempty"`
}

func (o ParseSqlLimitRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseSqlLimitRulesRequest struct{}"
	}

	return strings.Join([]string{"ParseSqlLimitRulesRequest", string(data)}, " ")
}

type ParseSqlLimitRulesRequestXLanguage struct {
	value string
}

type ParseSqlLimitRulesRequestXLanguageEnum struct {
	ZH_CN ParseSqlLimitRulesRequestXLanguage
	EN_US ParseSqlLimitRulesRequestXLanguage
}

func GetParseSqlLimitRulesRequestXLanguageEnum() ParseSqlLimitRulesRequestXLanguageEnum {
	return ParseSqlLimitRulesRequestXLanguageEnum{
		ZH_CN: ParseSqlLimitRulesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ParseSqlLimitRulesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ParseSqlLimitRulesRequestXLanguage) Value() string {
	return c.value
}

func (c ParseSqlLimitRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParseSqlLimitRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
