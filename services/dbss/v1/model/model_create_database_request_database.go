package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateDatabaseRequestDatabase 数据库信息
type CreateDatabaseRequestDatabase struct {

	// 数据库分类 - RDS: RDS数据库 - ECS:自建数据库
	DbClassification string `json:"db_classification"`

	// 数据库名称
	Name string `json:"name"`

	// 数据库类型 - MYSQL - ORACLE - POSTGRESQL - SQLSERVER - DAMENG - TAURUS - DWS - KINGBASE - GAUSSDBOPENGAUSS - GREENPLUM - HIGHGO - SHENTONG - GBASE8A - GBASE8S - GBASEXDM - MONGODB - DDS
	Type string `json:"type"`

	// 数据库版本
	Version string `json:"version"`

	// 字符集。 - GBK - UTF8
	Charset CreateDatabaseRequestDatabaseCharset `json:"charset"`

	// 数据库IP
	Ip string `json:"ip"`

	// 数据库端口
	Port string `json:"port"`

	// 数据库操作系统 - LINUX64 - WINDOWS64 - UNIX
	Os string `json:"os"`

	// 数据库实例名称
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o CreateDatabaseRequestDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequestDatabase struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequestDatabase", string(data)}, " ")
}

type CreateDatabaseRequestDatabaseCharset struct {
	value string
}

type CreateDatabaseRequestDatabaseCharsetEnum struct {
	GBK  CreateDatabaseRequestDatabaseCharset
	UTF8 CreateDatabaseRequestDatabaseCharset
}

func GetCreateDatabaseRequestDatabaseCharsetEnum() CreateDatabaseRequestDatabaseCharsetEnum {
	return CreateDatabaseRequestDatabaseCharsetEnum{
		GBK: CreateDatabaseRequestDatabaseCharset{
			value: "GBK",
		},
		UTF8: CreateDatabaseRequestDatabaseCharset{
			value: "UTF8",
		},
	}
}

func (c CreateDatabaseRequestDatabaseCharset) Value() string {
	return c.value
}

func (c CreateDatabaseRequestDatabaseCharset) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatabaseRequestDatabaseCharset) UnmarshalJSON(b []byte) error {
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
