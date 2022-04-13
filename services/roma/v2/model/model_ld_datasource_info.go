package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

type LdDatasourceInfo struct {
	// 数据源名称

	Name *string `json:"name,omitempty"`
	// 数据源类型： - oracle：oracle数据源类型 - mysql：mysql数据源类型 - mongodb：mongodb数据源类型 - redis：redis数据源类型 - postgresql：postgresql数据源类型 - hive：hive数据源类型 - mssql：sqlserver数据源类型 - sqlserver：sqlserver数据源类型 - gauss200：gauss200数据源类型 - dws：dws数据源类型 - gauss100：gauss100数据源类型 - zenith：zenith数据源类型

	Type *LdDatasourceInfoType `json:"type,omitempty"`
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
	// 数据源ID

	Id *string `json:"id,omitempty"`
	// 数据源状态： - old：存量数据源 - new：roma数据源

	Status *string `json:"status,omitempty"`
	// 数据源创建时间

	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`
	// 数据源更新时间

	ModifiedTime *sdktime.SdkTime `json:"modified_time,omitempty"`
}

func (o LdDatasourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdDatasourceInfo struct{}"
	}

	return strings.Join([]string{"LdDatasourceInfo", string(data)}, " ")
}

type LdDatasourceInfoType struct {
	value string
}

type LdDatasourceInfoTypeEnum struct {
	ORACLE     LdDatasourceInfoType
	MYSQL      LdDatasourceInfoType
	MONGODB    LdDatasourceInfoType
	REDIS      LdDatasourceInfoType
	POSTGRESQL LdDatasourceInfoType
	HIVE       LdDatasourceInfoType
	MSSQL      LdDatasourceInfoType
	SQLSERVER  LdDatasourceInfoType
	GAUSS200   LdDatasourceInfoType
	DWS        LdDatasourceInfoType
	GAUSS100   LdDatasourceInfoType
	ZENITH     LdDatasourceInfoType
}

func GetLdDatasourceInfoTypeEnum() LdDatasourceInfoTypeEnum {
	return LdDatasourceInfoTypeEnum{
		ORACLE: LdDatasourceInfoType{
			value: "oracle",
		},
		MYSQL: LdDatasourceInfoType{
			value: "mysql",
		},
		MONGODB: LdDatasourceInfoType{
			value: "mongodb",
		},
		REDIS: LdDatasourceInfoType{
			value: "redis",
		},
		POSTGRESQL: LdDatasourceInfoType{
			value: "postgresql",
		},
		HIVE: LdDatasourceInfoType{
			value: "hive",
		},
		MSSQL: LdDatasourceInfoType{
			value: "mssql",
		},
		SQLSERVER: LdDatasourceInfoType{
			value: "sqlserver",
		},
		GAUSS200: LdDatasourceInfoType{
			value: "gauss200",
		},
		DWS: LdDatasourceInfoType{
			value: "dws",
		},
		GAUSS100: LdDatasourceInfoType{
			value: "gauss100",
		},
		ZENITH: LdDatasourceInfoType{
			value: "zenith",
		},
	}
}

func (c LdDatasourceInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdDatasourceInfoType) UnmarshalJSON(b []byte) error {
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
