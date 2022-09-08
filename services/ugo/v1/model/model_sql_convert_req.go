package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SQL语句转换的请求体。
type SqlConvertReq struct {

	// 源数据库类型。
	SourceDbType SqlConvertReqSourceDbType `json:"source_db_type"`

	// 目标数据库类型。
	TargetDbType SqlConvertReqTargetDbType `json:"target_db_type"`

	// 目标数据库版本。 （注意：该字段需要与 target_db_type 字段组合成有效的目标数据库类型与版本，当前支持以下组合： GaussDB(for openGauss)-2020； RDS for PostgreSQL-11； RDS for PostgreSQL-Enhanced Edition； RDS for MySQL-5.7; GaussDB(for MySQL) 8.0。)
	TargetDbVersion SqlConvertReqTargetDbVersion `json:"target_db_version"`

	// 需要转换的SQL语句。
	SqlStatement string `json:"sql_statement"`
}

func (o SqlConvertReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlConvertReq struct{}"
	}

	return strings.Join([]string{"SqlConvertReq", string(data)}, " ")
}

type SqlConvertReqSourceDbType struct {
	value string
}

type SqlConvertReqSourceDbTypeEnum struct {
	ORACLE SqlConvertReqSourceDbType
}

func GetSqlConvertReqSourceDbTypeEnum() SqlConvertReqSourceDbTypeEnum {
	return SqlConvertReqSourceDbTypeEnum{
		ORACLE: SqlConvertReqSourceDbType{
			value: "ORACLE",
		},
	}
}

func (c SqlConvertReqSourceDbType) Value() string {
	return c.value
}

func (c SqlConvertReqSourceDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlConvertReqSourceDbType) UnmarshalJSON(b []byte) error {
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

type SqlConvertReqTargetDbType struct {
	value string
}

type SqlConvertReqTargetDbTypeEnum struct {
	RDS_FOR_MY_SQL          SqlConvertReqTargetDbType
	GAUSS_DB_FOR_OPEN_GAUSS SqlConvertReqTargetDbType
	GAUSS_DB_FOR_MY_SQL     SqlConvertReqTargetDbType
	RDS_FOR_POSTGRE_SQL     SqlConvertReqTargetDbType
}

func GetSqlConvertReqTargetDbTypeEnum() SqlConvertReqTargetDbTypeEnum {
	return SqlConvertReqTargetDbTypeEnum{
		RDS_FOR_MY_SQL: SqlConvertReqTargetDbType{
			value: "RDS for MySQL",
		},
		GAUSS_DB_FOR_OPEN_GAUSS: SqlConvertReqTargetDbType{
			value: "GaussDB(for openGauss)",
		},
		GAUSS_DB_FOR_MY_SQL: SqlConvertReqTargetDbType{
			value: "GaussDB(for MySQL)",
		},
		RDS_FOR_POSTGRE_SQL: SqlConvertReqTargetDbType{
			value: "RDS for PostgreSQL",
		},
	}
}

func (c SqlConvertReqTargetDbType) Value() string {
	return c.value
}

func (c SqlConvertReqTargetDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlConvertReqTargetDbType) UnmarshalJSON(b []byte) error {
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

type SqlConvertReqTargetDbVersion struct {
	value string
}

type SqlConvertReqTargetDbVersionEnum struct {
	E_5_7            SqlConvertReqTargetDbVersion
	E_8_0            SqlConvertReqTargetDbVersion
	E_11             SqlConvertReqTargetDbVersion
	E_2020           SqlConvertReqTargetDbVersion
	ENHANCED_EDITION SqlConvertReqTargetDbVersion
}

func GetSqlConvertReqTargetDbVersionEnum() SqlConvertReqTargetDbVersionEnum {
	return SqlConvertReqTargetDbVersionEnum{
		E_5_7: SqlConvertReqTargetDbVersion{
			value: "5.7",
		},
		E_8_0: SqlConvertReqTargetDbVersion{
			value: "8.0",
		},
		E_11: SqlConvertReqTargetDbVersion{
			value: "11",
		},
		E_2020: SqlConvertReqTargetDbVersion{
			value: "2020",
		},
		ENHANCED_EDITION: SqlConvertReqTargetDbVersion{
			value: "Enhanced Edition",
		},
	}
}

func (c SqlConvertReqTargetDbVersion) Value() string {
	return c.value
}

func (c SqlConvertReqTargetDbVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlConvertReqTargetDbVersion) UnmarshalJSON(b []byte) error {
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
