package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询迁移任务列表
type MigrationTaskList struct {

	// 迁移任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 迁移任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 迁移任务状态，这个字段的值包括：SUCCESS, FAILED, MIGRATING，TERMINATED
	Status *MigrationTaskListStatus `json:"status,omitempty"`

	// 迁移任务类型,包括备份文件导入和在线迁移两种类型。
	MigrationType *MigrationTaskListMigrationType `json:"migration_type,omitempty"`

	// 迁移方式，包括全量迁移和增量迁移两种类型。
	MigrationMethod *MigrationTaskListMigrationMethod `json:"migration_method,omitempty"`

	// 迁移机租户侧私有IP，与目的/源redis私有IP处于同VPC，可将此IP加入白名单
	EcsTenantPrivateIp *string `json:"ecs_tenant_private_ip,omitempty"`

	// 数据源，格式为ip:port或者桶名。
	DataSource *string `json:"data_source,omitempty"`

	// 源实例名称，若自建redis则为空。
	SourceInstanceName *string `json:"source_instance_name,omitempty"`

	// 源实例ID，若自建redis则为空。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// 目标redis地址，格式为ip:port。
	TargetInstanceAddrs *string `json:"target_instance_addrs,omitempty"`

	// 目标实例名称。
	TargetInstanceName *string `json:"target_instance_name,omitempty"`

	// 目标实例ID。
	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	// 迁移任务创建时间
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o MigrationTaskList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationTaskList struct{}"
	}

	return strings.Join([]string{"MigrationTaskList", string(data)}, " ")
}

type MigrationTaskListStatus struct {
	value string
}

type MigrationTaskListStatusEnum struct {
	SUCCESS    MigrationTaskListStatus
	FAILED     MigrationTaskListStatus
	MIGRATING  MigrationTaskListStatus
	TERMINATED MigrationTaskListStatus
}

func GetMigrationTaskListStatusEnum() MigrationTaskListStatusEnum {
	return MigrationTaskListStatusEnum{
		SUCCESS: MigrationTaskListStatus{
			value: "SUCCESS",
		},
		FAILED: MigrationTaskListStatus{
			value: "FAILED",
		},
		MIGRATING: MigrationTaskListStatus{
			value: "MIGRATING",
		},
		TERMINATED: MigrationTaskListStatus{
			value: "TERMINATED",
		},
	}
}

func (c MigrationTaskListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListStatus) UnmarshalJSON(b []byte) error {
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

type MigrationTaskListMigrationType struct {
	value string
}

type MigrationTaskListMigrationTypeEnum struct {
	BACKUPFILE_IMPORT MigrationTaskListMigrationType
	ONLINE_MIGRATION  MigrationTaskListMigrationType
}

func GetMigrationTaskListMigrationTypeEnum() MigrationTaskListMigrationTypeEnum {
	return MigrationTaskListMigrationTypeEnum{
		BACKUPFILE_IMPORT: MigrationTaskListMigrationType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: MigrationTaskListMigrationType{
			value: "online_migration",
		},
	}
}

func (c MigrationTaskListMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListMigrationType) UnmarshalJSON(b []byte) error {
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

type MigrationTaskListMigrationMethod struct {
	value string
}

type MigrationTaskListMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION MigrationTaskListMigrationMethod
	INCREMENTAL_MIGRATION MigrationTaskListMigrationMethod
}

func GetMigrationTaskListMigrationMethodEnum() MigrationTaskListMigrationMethodEnum {
	return MigrationTaskListMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: MigrationTaskListMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: MigrationTaskListMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c MigrationTaskListMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListMigrationMethod) UnmarshalJSON(b []byte) error {
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
