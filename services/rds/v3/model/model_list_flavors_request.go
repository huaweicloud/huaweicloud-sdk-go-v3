package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListFlavorsRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 数据库引擎。支持的引擎如下，不区分大小写： MySQL PostgreSQL SQLServer

	DatabaseName ListFlavorsRequestDatabaseName `json:"database_name"`
	// 数据库版本号，获取方法请参见5.1查询数据库引擎的版本。（可输入小版本号）

	VersionName *string `json:"version_name,omitempty"`
	// 规格编码

	SpecCode *string `json:"spec_code,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}

type ListFlavorsRequestDatabaseName struct {
	value string
}

type ListFlavorsRequestDatabaseNameEnum struct {
	MY_SQL      ListFlavorsRequestDatabaseName
	POSTGRE_SQL ListFlavorsRequestDatabaseName
	SQL_SERVER  ListFlavorsRequestDatabaseName
}

func GetListFlavorsRequestDatabaseNameEnum() ListFlavorsRequestDatabaseNameEnum {
	return ListFlavorsRequestDatabaseNameEnum{
		MY_SQL: ListFlavorsRequestDatabaseName{
			value: "MySQL",
		},
		POSTGRE_SQL: ListFlavorsRequestDatabaseName{
			value: "PostgreSQL",
		},
		SQL_SERVER: ListFlavorsRequestDatabaseName{
			value: "SQLServer",
		},
	}
}

func (c ListFlavorsRequestDatabaseName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorsRequestDatabaseName) UnmarshalJSON(b []byte) error {
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
