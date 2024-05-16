package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSqlLimitRulesBody 删除SQL限流规则请求体
type DeleteSqlLimitRulesBody struct {

	// 数据库类型
	DatastoreType DeleteSqlLimitRulesBodyDatastoreType `json:"datastore_type"`

	// SQL限流规则ID
	SqlLimitRuleIds []string `json:"sql_limit_rule_ids"`

	// 数据库名（PostgreSQL必填）
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o DeleteSqlLimitRulesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlLimitRulesBody struct{}"
	}

	return strings.Join([]string{"DeleteSqlLimitRulesBody", string(data)}, " ")
}

type DeleteSqlLimitRulesBodyDatastoreType struct {
	value string
}

type DeleteSqlLimitRulesBodyDatastoreTypeEnum struct {
	MY_SQL      DeleteSqlLimitRulesBodyDatastoreType
	POSTGRE_SQL DeleteSqlLimitRulesBodyDatastoreType
}

func GetDeleteSqlLimitRulesBodyDatastoreTypeEnum() DeleteSqlLimitRulesBodyDatastoreTypeEnum {
	return DeleteSqlLimitRulesBodyDatastoreTypeEnum{
		MY_SQL: DeleteSqlLimitRulesBodyDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: DeleteSqlLimitRulesBodyDatastoreType{
			value: "PostgreSQL",
		},
	}
}

func (c DeleteSqlLimitRulesBodyDatastoreType) Value() string {
	return c.value
}

func (c DeleteSqlLimitRulesBodyDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSqlLimitRulesBodyDatastoreType) UnmarshalJSON(b []byte) error {
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
