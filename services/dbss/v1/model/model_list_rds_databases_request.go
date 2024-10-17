package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRdsDatabasesRequest Request Object
type ListRdsDatabasesRequest struct {

	// 数据库类型 - MYSQL - ORACLE - POSTGRESQL - SQLSERVER - DAMENG - TAURUS - DWS - KINGBASE - MARIADB - GAUSSDBOPENGAUSS
	DbType ListRdsDatabasesRequestDbType `json:"db_type"`

	// 偏移量，从第一条数据偏移offset条数据后开始查询，默认为0。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数，默认为100。
	Limit *string `json:"limit,omitempty"`
}

func (o ListRdsDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRdsDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListRdsDatabasesRequest", string(data)}, " ")
}

type ListRdsDatabasesRequestDbType struct {
	value string
}

type ListRdsDatabasesRequestDbTypeEnum struct {
	MYSQL            ListRdsDatabasesRequestDbType
	ORACLE           ListRdsDatabasesRequestDbType
	POSTGRESQL       ListRdsDatabasesRequestDbType
	SQLSERVER        ListRdsDatabasesRequestDbType
	DAMENG           ListRdsDatabasesRequestDbType
	TAURUS           ListRdsDatabasesRequestDbType
	DWS              ListRdsDatabasesRequestDbType
	KINGBASE         ListRdsDatabasesRequestDbType
	MARIADB          ListRdsDatabasesRequestDbType
	GAUSSDBOPENGAUSS ListRdsDatabasesRequestDbType
}

func GetListRdsDatabasesRequestDbTypeEnum() ListRdsDatabasesRequestDbTypeEnum {
	return ListRdsDatabasesRequestDbTypeEnum{
		MYSQL: ListRdsDatabasesRequestDbType{
			value: "MYSQL",
		},
		ORACLE: ListRdsDatabasesRequestDbType{
			value: "ORACLE",
		},
		POSTGRESQL: ListRdsDatabasesRequestDbType{
			value: "POSTGRESQL",
		},
		SQLSERVER: ListRdsDatabasesRequestDbType{
			value: "SQLSERVER",
		},
		DAMENG: ListRdsDatabasesRequestDbType{
			value: "DAMENG",
		},
		TAURUS: ListRdsDatabasesRequestDbType{
			value: "TAURUS",
		},
		DWS: ListRdsDatabasesRequestDbType{
			value: "DWS",
		},
		KINGBASE: ListRdsDatabasesRequestDbType{
			value: "KINGBASE",
		},
		MARIADB: ListRdsDatabasesRequestDbType{
			value: "MARIADB",
		},
		GAUSSDBOPENGAUSS: ListRdsDatabasesRequestDbType{
			value: "GAUSSDBOPENGAUSS",
		},
	}
}

func (c ListRdsDatabasesRequestDbType) Value() string {
	return c.value
}

func (c ListRdsDatabasesRequestDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRdsDatabasesRequestDbType) UnmarshalJSON(b []byte) error {
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
