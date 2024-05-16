package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSqlLimitRulesBody 修改SQL限流规则请求体
type UpdateSqlLimitRulesBody struct {

	// 数据库类型
	DatastoreType UpdateSqlLimitRulesBodyDatastoreType `json:"datastore_type"`

	// SQL限流规则ID
	SqlLimitRuleIds []string `json:"sql_limit_rule_ids"`

	// 数据库名（PostgreSQL必填）
	DatabaseName *string `json:"database_name,omitempty"`

	SqlLimitRule *UpdateSqlLimitRuleOption `json:"sql_limit_rule"`
}

func (o UpdateSqlLimitRulesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSqlLimitRulesBody struct{}"
	}

	return strings.Join([]string{"UpdateSqlLimitRulesBody", string(data)}, " ")
}

type UpdateSqlLimitRulesBodyDatastoreType struct {
	value string
}

type UpdateSqlLimitRulesBodyDatastoreTypeEnum struct {
	POSTGRE_SQL UpdateSqlLimitRulesBodyDatastoreType
}

func GetUpdateSqlLimitRulesBodyDatastoreTypeEnum() UpdateSqlLimitRulesBodyDatastoreTypeEnum {
	return UpdateSqlLimitRulesBodyDatastoreTypeEnum{
		POSTGRE_SQL: UpdateSqlLimitRulesBodyDatastoreType{
			value: "PostgreSQL",
		},
	}
}

func (c UpdateSqlLimitRulesBodyDatastoreType) Value() string {
	return c.value
}

func (c UpdateSqlLimitRulesBodyDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSqlLimitRulesBodyDatastoreType) UnmarshalJSON(b []byte) error {
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
