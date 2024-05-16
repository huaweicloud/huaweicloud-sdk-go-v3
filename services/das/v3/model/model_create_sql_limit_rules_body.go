package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSqlLimitRulesBody 创建SQL限流规则请求提
type CreateSqlLimitRulesBody struct {

	// 数据库类型
	DatastoreType CreateSqlLimitRulesBodyDatastoreType `json:"datastore_type"`

	// 需要创建的SQL限流规则列表，一次最多创建5个
	SqlLimitRules []CreateSqlLimitRuleOption `json:"sql_limit_rules"`

	// 数据库名（PostgreSQL必填）
	DatabaseName *string `json:"database_name,omitempty"`
}

func (o CreateSqlLimitRulesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitRulesBody struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitRulesBody", string(data)}, " ")
}

type CreateSqlLimitRulesBodyDatastoreType struct {
	value string
}

type CreateSqlLimitRulesBodyDatastoreTypeEnum struct {
	MY_SQL      CreateSqlLimitRulesBodyDatastoreType
	POSTGRE_SQL CreateSqlLimitRulesBodyDatastoreType
}

func GetCreateSqlLimitRulesBodyDatastoreTypeEnum() CreateSqlLimitRulesBodyDatastoreTypeEnum {
	return CreateSqlLimitRulesBodyDatastoreTypeEnum{
		MY_SQL: CreateSqlLimitRulesBodyDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: CreateSqlLimitRulesBodyDatastoreType{
			value: "PostgreSQL",
		},
	}
}

func (c CreateSqlLimitRulesBodyDatastoreType) Value() string {
	return c.value
}

func (c CreateSqlLimitRulesBodyDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSqlLimitRulesBodyDatastoreType) UnmarshalJSON(b []byte) error {
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
