package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListMigrationTaskResponse struct {

	// 迁移任务数量。
	Count *int32 `json:"count,omitempty"`

	// 迁移任务列表。
	MigrationTasks *[]MigrationTaskList `json:"migration_tasks,omitempty"`

	// 目标实例地址
	TargetInstanceAddress *string `json:"target_instance_address,omitempty"`

	// 迁移方式，包括全量迁移和增量迁移两种类型。
	MigrationMethod *ListMigrationTaskResponseMigrationMethod `json:"migration_method,omitempty"`

	// 迁移任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 目标实例ID。
	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	// 源实例名称，若自建redis则为空。
	SourceInstanceName *string `json:"source_instance_name,omitempty"`

	// 目标实例名称。
	TargetInstanceName *string `json:"target_instance_name,omitempty"`

	// 迁移任务类型,包括备份文件导入和在线迁移两种类型。
	MigrateType *ListMigrationTaskResponseMigrateType `json:"migrate_type,omitempty"`

	// 迁移任务创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 源实例ID，若自建redis则为空。
	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	// 迁移任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 源redis地址，格式为ip:port或者桶名。
	DataSource *string `json:"data_source,omitempty"`

	// 迁移任务状态，这个字段的值包括：SUCCESS, FAILED, MIGRATING，TERMINATED
	Status         *ListMigrationTaskResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskResponse", string(data)}, " ")
}

type ListMigrationTaskResponseMigrationMethod struct {
	value string
}

type ListMigrationTaskResponseMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION ListMigrationTaskResponseMigrationMethod
	INCREMENTAL_MIGRATION ListMigrationTaskResponseMigrationMethod
}

func GetListMigrationTaskResponseMigrationMethodEnum() ListMigrationTaskResponseMigrationMethodEnum {
	return ListMigrationTaskResponseMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: ListMigrationTaskResponseMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: ListMigrationTaskResponseMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c ListMigrationTaskResponseMigrationMethod) Value() string {
	return c.value
}

func (c ListMigrationTaskResponseMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMigrationTaskResponseMigrationMethod) UnmarshalJSON(b []byte) error {
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

type ListMigrationTaskResponseMigrateType struct {
	value string
}

type ListMigrationTaskResponseMigrateTypeEnum struct {
	BACKUPFILE_IMPORT ListMigrationTaskResponseMigrateType
	ONLINE_MIGRATION  ListMigrationTaskResponseMigrateType
}

func GetListMigrationTaskResponseMigrateTypeEnum() ListMigrationTaskResponseMigrateTypeEnum {
	return ListMigrationTaskResponseMigrateTypeEnum{
		BACKUPFILE_IMPORT: ListMigrationTaskResponseMigrateType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: ListMigrationTaskResponseMigrateType{
			value: "online_migration",
		},
	}
}

func (c ListMigrationTaskResponseMigrateType) Value() string {
	return c.value
}

func (c ListMigrationTaskResponseMigrateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMigrationTaskResponseMigrateType) UnmarshalJSON(b []byte) error {
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

type ListMigrationTaskResponseStatus struct {
	value string
}

type ListMigrationTaskResponseStatusEnum struct {
	SUCCESS    ListMigrationTaskResponseStatus
	FAILED     ListMigrationTaskResponseStatus
	MIGRATING  ListMigrationTaskResponseStatus
	TERMINATED ListMigrationTaskResponseStatus
}

func GetListMigrationTaskResponseStatusEnum() ListMigrationTaskResponseStatusEnum {
	return ListMigrationTaskResponseStatusEnum{
		SUCCESS: ListMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: ListMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: ListMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: ListMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c ListMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c ListMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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
