package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LdDatasourceCreate struct {

	// 数据源名称
	Name *string `json:"name,omitempty"`

	// 数据源类型： - oracle：oracle数据源类型 - mysql：mysql数据源类型 - mongodb：mongodb数据源类型 - redis：redis数据源类型 - postgresql：postgresql数据源类型 - hive：hive数据源类型 - mssql：sqlserver数据源类型 - sqlserver：sqlserver数据源类型 - gauss200：gauss200数据源类型 - dws：dws数据源类型 - gauss100：gauss100数据源类型 - zenith：zenith数据源类型
	Type *LdDatasourceCreateType `json:"type,omitempty"`

	// 数据源描述
	Description *string `json:"description,omitempty"`

	// 数据源连接字符串
	Url *string `json:"url,omitempty"`

	// 用户名
	User *string `json:"user,omitempty"`

	// 密码。  敏感信息不作为响应返回
	Password *string `json:"password,omitempty"`

	// ftp上传路径  预留字段，暂不支持。
	Remotepath *string `json:"remotepath,omitempty"`
}

func (o LdDatasourceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdDatasourceCreate struct{}"
	}

	return strings.Join([]string{"LdDatasourceCreate", string(data)}, " ")
}

type LdDatasourceCreateType struct {
	value string
}

type LdDatasourceCreateTypeEnum struct {
	ORACLE     LdDatasourceCreateType
	MYSQL      LdDatasourceCreateType
	MONGODB    LdDatasourceCreateType
	REDIS      LdDatasourceCreateType
	POSTGRESQL LdDatasourceCreateType
	HIVE       LdDatasourceCreateType
	MSSQL      LdDatasourceCreateType
	SQLSERVER  LdDatasourceCreateType
	GAUSS200   LdDatasourceCreateType
	DWS        LdDatasourceCreateType
	GAUSS100   LdDatasourceCreateType
	ZENITH     LdDatasourceCreateType
}

func GetLdDatasourceCreateTypeEnum() LdDatasourceCreateTypeEnum {
	return LdDatasourceCreateTypeEnum{
		ORACLE: LdDatasourceCreateType{
			value: "oracle",
		},
		MYSQL: LdDatasourceCreateType{
			value: "mysql",
		},
		MONGODB: LdDatasourceCreateType{
			value: "mongodb",
		},
		REDIS: LdDatasourceCreateType{
			value: "redis",
		},
		POSTGRESQL: LdDatasourceCreateType{
			value: "postgresql",
		},
		HIVE: LdDatasourceCreateType{
			value: "hive",
		},
		MSSQL: LdDatasourceCreateType{
			value: "mssql",
		},
		SQLSERVER: LdDatasourceCreateType{
			value: "sqlserver",
		},
		GAUSS200: LdDatasourceCreateType{
			value: "gauss200",
		},
		DWS: LdDatasourceCreateType{
			value: "dws",
		},
		GAUSS100: LdDatasourceCreateType{
			value: "gauss100",
		},
		ZENITH: LdDatasourceCreateType{
			value: "zenith",
		},
	}
}

func (c LdDatasourceCreateType) Value() string {
	return c.value
}

func (c LdDatasourceCreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdDatasourceCreateType) UnmarshalJSON(b []byte) error {
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
