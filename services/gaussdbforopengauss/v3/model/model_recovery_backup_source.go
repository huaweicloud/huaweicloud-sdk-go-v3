package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RecoveryBackupSource struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 恢复备份类型：backup，timestamp，different
	Type *string `json:"type,omitempty"`

	// 用于恢复的备份ID。
	BackupId string `json:"backup_id"`

	// UTC时间，时间戳
	RestoreTime string `json:"restore_time"`

	// 表基础信息。
	TableList *[]RestoreTableListDetail `json:"table_list,omitempty"`

	// 备份级别取值, 默认值：INSTANCE
	SchemaType *RecoveryBackupSourceSchemaType `json:"schema_type,omitempty"`
}

func (o RecoveryBackupSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecoveryBackupSource struct{}"
	}

	return strings.Join([]string{"RecoveryBackupSource", string(data)}, " ")
}

type RecoveryBackupSourceSchemaType struct {
	value string
}

type RecoveryBackupSourceSchemaTypeEnum struct {
	INSTANCE       RecoveryBackupSourceSchemaType
	DATABASE       RecoveryBackupSourceSchemaType
	DATABASE_TABLE RecoveryBackupSourceSchemaType
}

func GetRecoveryBackupSourceSchemaTypeEnum() RecoveryBackupSourceSchemaTypeEnum {
	return RecoveryBackupSourceSchemaTypeEnum{
		INSTANCE: RecoveryBackupSourceSchemaType{
			value: "INSTANCE 实例级备份",
		},
		DATABASE: RecoveryBackupSourceSchemaType{
			value: "DATABASE 库级备份",
		},
		DATABASE_TABLE: RecoveryBackupSourceSchemaType{
			value: "DATABASE_TABLE 表级备份",
		},
	}
}

func (c RecoveryBackupSourceSchemaType) Value() string {
	return c.value
}

func (c RecoveryBackupSourceSchemaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecoveryBackupSourceSchemaType) UnmarshalJSON(b []byte) error {
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
