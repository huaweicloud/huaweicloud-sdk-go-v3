package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type BackupDatastore struct {
	// 数据库引擎，不区分大小写：  - MySQL - PostgreSQL - SQLServer

	Type BackupDatastoreType `json:"type"`
	// 数据库版本。

	Version string `json:"version"`
}

func (o BackupDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupDatastore struct{}"
	}

	return strings.Join([]string{"BackupDatastore", string(data)}, " ")
}

type BackupDatastoreType struct {
	value string
}

type BackupDatastoreTypeEnum struct {
	MY_SQL      BackupDatastoreType
	POSTGRE_SQL BackupDatastoreType
	SQL_SERVER  BackupDatastoreType
}

func GetBackupDatastoreTypeEnum() BackupDatastoreTypeEnum {
	return BackupDatastoreTypeEnum{
		MY_SQL: BackupDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: BackupDatastoreType{
			value: "PostgreSQL",
		},
		SQL_SERVER: BackupDatastoreType{
			value: "SQLServer",
		},
	}
}

func (c BackupDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupDatastoreType) UnmarshalJSON(b []byte) error {
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
