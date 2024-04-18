package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StopMigrationTaskResponse Response Object
type StopMigrationTaskResponse struct {

	// 迁移任务ID。
	TaskId *string `json:"task_id,omitempty"`

	// 迁移任务名称。
	TaskName *string `json:"task_name,omitempty"`

	// 迁移任务描述。
	Description *string `json:"description,omitempty"`

	// 迁移任务状态，这个字段的值包括：SUCCESS, FAILED, MIGRATING，TERMINATED。
	Status *StopMigrationTaskResponseStatus `json:"status,omitempty"`

	// 迁移任务类型,包括备份文件导入和在线迁移两种类型。
	MigrationType *StopMigrationTaskResponseMigrationType `json:"migration_type,omitempty"`

	// 迁移方式，包括全量迁移和增量迁移两种类型。
	MigrationMethod *StopMigrationTaskResponseMigrationMethod `json:"migration_method,omitempty"`

	// 迁移机租户侧私有IP，与目的/源redis私有IP处于同VPC，可将此IP加入白名单
	EcsTenantPrivateIp *string `json:"ecs_tenant_private_ip,omitempty"`

	BackupFiles *BackupFilesBody `json:"backup_files,omitempty"`

	// 网络类型，包括vpc和vpn两种类型。
	NetworkType *StopMigrationTaskResponseNetworkType `json:"network_type,omitempty"`

	SourceInstance *SourceInstanceBody `json:"source_instance,omitempty"`

	TargetInstance *TargetInstanceBody `json:"target_instance,omitempty"`

	// 迁移任务创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 迁移任务完成时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 迁移机释放时间。
	ReleasedAt *string `json:"released_at,omitempty"`

	// 版本。
	Version *string `json:"version,omitempty"`

	// 操作模式，分为auto和manual。
	ResumeMode *string `json:"resume_mode,omitempty"`

	// 支持的特性。
	SupportedFeatures *[]string `json:"supported_features,omitempty"`

	// 租户VPC ID。
	TenantVpcId *string `json:"tenant_vpc_id,omitempty"`

	// 租户子网ID。
	TenantSubnetId *string `json:"tenant_subnet_id,omitempty"`

	// 租户安全组ID。
	TenantSecurityGroupId *string `json:"tenant_security_group_id,omitempty"`

	// 带宽限制速度。
	BandwidthLimitMb *string `json:"bandwidth_limit_mb,omitempty"`

	// 任务状态。
	TaskStatus     *string `json:"task_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskResponse", string(data)}, " ")
}

type StopMigrationTaskResponseStatus struct {
	value string
}

type StopMigrationTaskResponseStatusEnum struct {
	SUCCESS    StopMigrationTaskResponseStatus
	FAILED     StopMigrationTaskResponseStatus
	MIGRATING  StopMigrationTaskResponseStatus
	TERMINATED StopMigrationTaskResponseStatus
}

func GetStopMigrationTaskResponseStatusEnum() StopMigrationTaskResponseStatusEnum {
	return StopMigrationTaskResponseStatusEnum{
		SUCCESS: StopMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: StopMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: StopMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: StopMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c StopMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type StopMigrationTaskResponseMigrationType struct {
	value string
}

type StopMigrationTaskResponseMigrationTypeEnum struct {
	BACKUPFILE_IMPORT StopMigrationTaskResponseMigrationType
	ONLINE_MIGRATION  StopMigrationTaskResponseMigrationType
}

func GetStopMigrationTaskResponseMigrationTypeEnum() StopMigrationTaskResponseMigrationTypeEnum {
	return StopMigrationTaskResponseMigrationTypeEnum{
		BACKUPFILE_IMPORT: StopMigrationTaskResponseMigrationType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: StopMigrationTaskResponseMigrationType{
			value: "online_migration",
		},
	}
}

func (c StopMigrationTaskResponseMigrationType) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseMigrationType) UnmarshalJSON(b []byte) error {
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

type StopMigrationTaskResponseMigrationMethod struct {
	value string
}

type StopMigrationTaskResponseMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION StopMigrationTaskResponseMigrationMethod
	INCREMENTAL_MIGRATION StopMigrationTaskResponseMigrationMethod
}

func GetStopMigrationTaskResponseMigrationMethodEnum() StopMigrationTaskResponseMigrationMethodEnum {
	return StopMigrationTaskResponseMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: StopMigrationTaskResponseMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: StopMigrationTaskResponseMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c StopMigrationTaskResponseMigrationMethod) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseMigrationMethod) UnmarshalJSON(b []byte) error {
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

type StopMigrationTaskResponseNetworkType struct {
	value string
}

type StopMigrationTaskResponseNetworkTypeEnum struct {
	VPC StopMigrationTaskResponseNetworkType
	VPN StopMigrationTaskResponseNetworkType
}

func GetStopMigrationTaskResponseNetworkTypeEnum() StopMigrationTaskResponseNetworkTypeEnum {
	return StopMigrationTaskResponseNetworkTypeEnum{
		VPC: StopMigrationTaskResponseNetworkType{
			value: "vpc",
		},
		VPN: StopMigrationTaskResponseNetworkType{
			value: "vpn",
		},
	}
}

func (c StopMigrationTaskResponseNetworkType) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseNetworkType) UnmarshalJSON(b []byte) error {
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
