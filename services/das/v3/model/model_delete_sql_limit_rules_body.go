package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 删除SQL限流规则请求体
type DeleteSqlLimitRulesBody struct {

	// 数据库类型
	DatastoreType DeleteSqlLimitRulesBodyDatastoreType `json:"datastore_type"`

	// SQL限流规则ID
	SqlLimitRuleIds []string `json:"sql_limit_rule_ids"`
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
	MY_SQL DeleteSqlLimitRulesBodyDatastoreType
}

func GetDeleteSqlLimitRulesBodyDatastoreTypeEnum() DeleteSqlLimitRulesBodyDatastoreTypeEnum {
	return DeleteSqlLimitRulesBodyDatastoreTypeEnum{
		MY_SQL: DeleteSqlLimitRulesBodyDatastoreType{
			value: "MySQL",
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
