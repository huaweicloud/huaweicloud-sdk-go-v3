package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResourceTypeEnum Server：云服务器 Volume：云硬盘 Sfs-Turbo：高性能文件系统 Workspace：云桌面 MySQL：云数据库RDS(MySQL) PostgreSQL：云数据库RDS(PostgreSQL) SQLServer：云数据库RDS(SQLServer) MariaDB：云数据库RDS(MariaDB) GaussDB：云数据库GaussDB
type ResourceTypeEnum struct {
	value string
}

type ResourceTypeEnumEnum struct {
	SERVER      ResourceTypeEnum
	VOLUME      ResourceTypeEnum
	SFS_TURBO   ResourceTypeEnum
	WORKSPACE   ResourceTypeEnum
	MY_SQL      ResourceTypeEnum
	POSTGRE_SQL ResourceTypeEnum
	SQL_SERVER  ResourceTypeEnum
	MARIA_DB    ResourceTypeEnum
	GAUSS_DB    ResourceTypeEnum
}

func GetResourceTypeEnumEnum() ResourceTypeEnumEnum {
	return ResourceTypeEnumEnum{
		SERVER: ResourceTypeEnum{
			value: "Server",
		},
		VOLUME: ResourceTypeEnum{
			value: "Volume",
		},
		SFS_TURBO: ResourceTypeEnum{
			value: "Sfs-Turbo",
		},
		WORKSPACE: ResourceTypeEnum{
			value: "Workspace",
		},
		MY_SQL: ResourceTypeEnum{
			value: "MySQL",
		},
		POSTGRE_SQL: ResourceTypeEnum{
			value: "PostgreSQL",
		},
		SQL_SERVER: ResourceTypeEnum{
			value: "SQLServer",
		},
		MARIA_DB: ResourceTypeEnum{
			value: "MariaDB",
		},
		GAUSS_DB: ResourceTypeEnum{
			value: "GaussDB",
		},
	}
}

func (c ResourceTypeEnum) Value() string {
	return c.value
}

func (c ResourceTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceTypeEnum) UnmarshalJSON(b []byte) error {
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
