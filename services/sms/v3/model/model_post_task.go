package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostTask 创建任务的参数
type PostTask struct {

	// 任务名称
	Name string `json:"name"`

	// 任务类型 MIGRATE_FILE:文件级迁移 MIGRATE_BLOCK:块级迁移
	Type PostTaskType `json:"type"`

	// 迁移后是否启动目的端虚拟机
	StartTargetServer *bool `json:"start_target_server,omitempty"`

	// 是否自动启动
	AutoStart *bool `json:"auto_start,omitempty"`

	// 操作系统类型 WINDOWS:Windows系统类型 LINUX:Linux系统类型
	OsType PostTaskOsType `json:"os_type"`

	SourceServer *SourceServerByTask `json:"source_server"`

	TargetServer *TargetServerByTask `json:"target_server"`

	// 迁移IP，如果是自动创建虚拟机，不需要此参数
	MigrationIp *string `json:"migration_ip,omitempty"`

	// region的名称
	RegionName string `json:"region_name"`

	// region ID
	RegionId string `json:"region_id"`

	// 项目名称
	ProjectName string `json:"project_name"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 优先级。默认为1
	Priority *int32 `json:"priority,omitempty"`

	// 自动创建虚拟机使用模板
	VmTemplateId *string `json:"vm_template_id,omitempty"`

	// 是否使用公网ip
	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	// 是否使用ipv6
	UseIpv6 *bool `json:"use_ipv6,omitempty"`

	// 复制或者同步后是否会继续持续同步，不添加则默认是false
	Syncing *bool `json:"syncing,omitempty"`

	// 是否存在服务，如果存在，则创建任务
	ExistServer *bool `json:"exist_server,omitempty"`

	// 是否开启网络检测
	StartNetworkCheck *bool `json:"start_network_check,omitempty"`

	// 迁移速率限制值
	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	// 停止迁移的超速阈值。 是一个迁移速率的保护机制，超出该阈值会停止任务。它主要用于控制迁移过程中资源（特别是网络带宽）的消耗，确保系统的整体性能不受单一迁移任务影响 单位是百分比
	OverSpeedThreshold *float64 `json:"over_speed_threshold,omitempty"`

	// 是否进行一致性校验
	IsNeedConsistencyCheck *bool `json:"is_need_consistency_check,omitempty"`

	// 是否开启迁移演练
	NeedMigrationTest *bool `json:"need_migration_test,omitempty"`
}

func (o PostTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostTask struct{}"
	}

	return strings.Join([]string{"PostTask", string(data)}, " ")
}

type PostTaskType struct {
	value string
}

type PostTaskTypeEnum struct {
	MIGRATE_FILE  PostTaskType
	MIGRATE_BLOCK PostTaskType
}

func GetPostTaskTypeEnum() PostTaskTypeEnum {
	return PostTaskTypeEnum{
		MIGRATE_FILE: PostTaskType{
			value: "MIGRATE_FILE",
		},
		MIGRATE_BLOCK: PostTaskType{
			value: "MIGRATE_BLOCK",
		},
	}
}

func (c PostTaskType) Value() string {
	return c.value
}

func (c PostTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostTaskType) UnmarshalJSON(b []byte) error {
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

type PostTaskOsType struct {
	value string
}

type PostTaskOsTypeEnum struct {
	WINDOWS PostTaskOsType
	LINUX   PostTaskOsType
}

func GetPostTaskOsTypeEnum() PostTaskOsTypeEnum {
	return PostTaskOsTypeEnum{
		WINDOWS: PostTaskOsType{
			value: "WINDOWS",
		},
		LINUX: PostTaskOsType{
			value: "LINUX",
		},
	}
}

func (c PostTaskOsType) Value() string {
	return c.value
}

func (c PostTaskOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostTaskOsType) UnmarshalJSON(b []byte) error {
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
