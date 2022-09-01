package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowMigrationTaskResponse struct {

	// 迁移任务ID。
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 迁移任务名称。
	TaskName *string `json:"task_name,omitempty" xml:"task_name"`

	// 迁移任务描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 迁移任务状态，这个字段的值包括：SUCCESS, FAILED, MIGRATING，TERMINATED。
	Status *ShowMigrationTaskResponseStatus `json:"status,omitempty" xml:"status"`

	// 迁移任务类型,包括备份文件导入和在线迁移两种类型。
	MigrationType *ShowMigrationTaskResponseMigrationType `json:"migration_type,omitempty" xml:"migration_type"`

	// 迁移方式，包括全量迁移和增量迁移两种类型。
	MigrationMethod *ShowMigrationTaskResponseMigrationMethod `json:"migration_method,omitempty" xml:"migration_method"`

	// 迁移机租户侧私有IP，与目的/源redis私有IP处于同VPC，可将此IP加入白名单
	EcsTenantPrivateIp *string `json:"ecs_tenant_private_ip,omitempty" xml:"ecs_tenant_private_ip"`

	BackupFiles *BackupFilesBody `json:"backup_files,omitempty" xml:"backup_files"`

	// 网络类型，包括vpc和vpn两种类型。
	NetworkType *ShowMigrationTaskResponseNetworkType `json:"network_type,omitempty" xml:"network_type"`

	SourceInstance *SourceInstanceBody `json:"source_instance,omitempty" xml:"source_instance"`

	TargetInstance *TargetInstanceBody `json:"target_instance,omitempty" xml:"target_instance"`

	// 迁移任务创建时间。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 迁移任务完成时间。
	UpdatedAt      *string `json:"updated_at,omitempty" xml:"updated_at"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskResponse", string(data)}, " ")
}

type ShowMigrationTaskResponseStatus struct {
	value string
}

type ShowMigrationTaskResponseStatusEnum struct {
	SUCCESS    ShowMigrationTaskResponseStatus
	FAILED     ShowMigrationTaskResponseStatus
	MIGRATING  ShowMigrationTaskResponseStatus
	TERMINATED ShowMigrationTaskResponseStatus
}

func GetShowMigrationTaskResponseStatusEnum() ShowMigrationTaskResponseStatusEnum {
	return ShowMigrationTaskResponseStatusEnum{
		SUCCESS: ShowMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: ShowMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: ShowMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: ShowMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c ShowMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowMigrationTaskResponseMigrationType struct {
	value string
}

type ShowMigrationTaskResponseMigrationTypeEnum struct {
	BACKUPFILE_IMPORT ShowMigrationTaskResponseMigrationType
	ONLINE_MIGRATION  ShowMigrationTaskResponseMigrationType
}

func GetShowMigrationTaskResponseMigrationTypeEnum() ShowMigrationTaskResponseMigrationTypeEnum {
	return ShowMigrationTaskResponseMigrationTypeEnum{
		BACKUPFILE_IMPORT: ShowMigrationTaskResponseMigrationType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: ShowMigrationTaskResponseMigrationType{
			value: "online_migration",
		},
	}
}

func (c ShowMigrationTaskResponseMigrationType) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseMigrationType) UnmarshalJSON(b []byte) error {
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

type ShowMigrationTaskResponseMigrationMethod struct {
	value string
}

type ShowMigrationTaskResponseMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION ShowMigrationTaskResponseMigrationMethod
	INCREMENTAL_MIGRATION ShowMigrationTaskResponseMigrationMethod
}

func GetShowMigrationTaskResponseMigrationMethodEnum() ShowMigrationTaskResponseMigrationMethodEnum {
	return ShowMigrationTaskResponseMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: ShowMigrationTaskResponseMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: ShowMigrationTaskResponseMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c ShowMigrationTaskResponseMigrationMethod) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseMigrationMethod) UnmarshalJSON(b []byte) error {
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

type ShowMigrationTaskResponseNetworkType struct {
	value string
}

type ShowMigrationTaskResponseNetworkTypeEnum struct {
	VPC ShowMigrationTaskResponseNetworkType
	VPN ShowMigrationTaskResponseNetworkType
}

func GetShowMigrationTaskResponseNetworkTypeEnum() ShowMigrationTaskResponseNetworkTypeEnum {
	return ShowMigrationTaskResponseNetworkTypeEnum{
		VPC: ShowMigrationTaskResponseNetworkType{
			value: "vpc",
		},
		VPN: ShowMigrationTaskResponseNetworkType{
			value: "vpn",
		},
	}
}

func (c ShowMigrationTaskResponseNetworkType) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseNetworkType) UnmarshalJSON(b []byte) error {
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
