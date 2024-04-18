package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MigrationTaskList 查询迁移任务列表
type MigrationTaskList struct {

	// 迁移任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 迁移任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 迁移任务状态，这个字段的值包括：SUCCESS（成功）, FAILED（失败）, MIGRATING（迁移中），TERMINATED（已结束）。
	Status *MigrationTaskListStatus `json:"status,omitempty"`

	// 迁移任务类型，包括备份文件导入和在线迁移两种类型。
	MigrationType *MigrationTaskListMigrationType `json:"migration_type,omitempty"`

	// 迁移方式，包括全量迁移和增量迁移两种类型。
	MigrationMethod *MigrationTaskListMigrationMethod `json:"migration_method,omitempty"`

	// 迁移机租户侧私有IP，与目的/源redis私有IP处于同VPC，可将此IP加入白名单。
	EcsTenantPrivateIp *string `json:"ecs_tenant_private_ip,omitempty"`

	// 源redis地址，格式为ip:port或者桶名。
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

	// 迁移任务创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 迁移任务描述。
	Description *string `json:"description,omitempty"`

	// 源实例状态，若自建redis则为空。
	SourceInstanceStatus *string `json:"source_instance_status,omitempty"`

	// 目标实例状态。
	TargetInstanceStatus *string `json:"target_instance_status,omitempty"`

	// 源实例子网ID，若自建redis则为空。
	SourceInstanceSubnetId *string `json:"source_instance_subnet_id,omitempty"`

	// 目标实例子网ID。
	TargetInstanceSubnetId *string `json:"target_instance_subnet_id,omitempty"`

	// 源实例规格编码，若自建redis则为空。
	SourceInstanceSpecCode *string `json:"source_instance_spec_code,omitempty"`

	// 目标实例规格编码。
	TargetInstanceSpecCode *string `json:"target_instance_spec_code,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 迁移机释放时间。
	ReleasedAt *string `json:"released_at,omitempty"`

	// 版本。
	Version *string `json:"version,omitempty"`

	// 操作模式，分为auto和manual。
	ResumeMode *string `json:"resume_mode,omitempty"`

	// 支持的特性。
	SupportedFeatures *[]string `json:"supported_features,omitempty"`
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

func (c MigrationTaskListStatus) Value() string {
	return c.value
}

func (c MigrationTaskListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListStatus) UnmarshalJSON(b []byte) error {
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

func (c MigrationTaskListMigrationType) Value() string {
	return c.value
}

func (c MigrationTaskListMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListMigrationType) UnmarshalJSON(b []byte) error {
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

func (c MigrationTaskListMigrationMethod) Value() string {
	return c.value
}

func (c MigrationTaskListMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListMigrationMethod) UnmarshalJSON(b []byte) error {
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
