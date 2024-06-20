package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ParseSqlLimitRulesReq {\"originalSql\":\"select * from das_conn_info\",\"useTemplate\":true,\"keepOperators\":true}
type ParseSqlLimitRulesReq struct {

	// 数据库类型，目前支持MySQL、MariaDB、GaussDB(for MySQL)三种引擎。
	DatastoreType ParseSqlLimitRulesReqDatastoreType `json:"datastore_type"`

	// 原始SQL语句
	OriginalSql string `json:"original_sql"`

	// 是否校验SQL语句
	UseTemplate bool `json:"use_template"`

	// 是否保留操作符
	KeepOperators bool `json:"keep_operators"`
}

func (o ParseSqlLimitRulesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseSqlLimitRulesReq struct{}"
	}

	return strings.Join([]string{"ParseSqlLimitRulesReq", string(data)}, " ")
}

type ParseSqlLimitRulesReqDatastoreType struct {
	value string
}

type ParseSqlLimitRulesReqDatastoreTypeEnum struct {
	MY_SQL              ParseSqlLimitRulesReqDatastoreType
	MARIA_DB            ParseSqlLimitRulesReqDatastoreType
	GAUSS_DB_FOR_MY_SQL ParseSqlLimitRulesReqDatastoreType
}

func GetParseSqlLimitRulesReqDatastoreTypeEnum() ParseSqlLimitRulesReqDatastoreTypeEnum {
	return ParseSqlLimitRulesReqDatastoreTypeEnum{
		MY_SQL: ParseSqlLimitRulesReqDatastoreType{
			value: "MySQL",
		},
		MARIA_DB: ParseSqlLimitRulesReqDatastoreType{
			value: "MariaDB",
		},
		GAUSS_DB_FOR_MY_SQL: ParseSqlLimitRulesReqDatastoreType{
			value: "GaussDB(for MySQL)",
		},
	}
}

func (c ParseSqlLimitRulesReqDatastoreType) Value() string {
	return c.value
}

func (c ParseSqlLimitRulesReqDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ParseSqlLimitRulesReqDatastoreType) UnmarshalJSON(b []byte) error {
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
