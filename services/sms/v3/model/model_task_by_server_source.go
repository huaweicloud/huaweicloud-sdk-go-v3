package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskByServerSource 源端列表中关联的任务
type TaskByServerSource struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 迁移项目类型 MIGRATE_BLOCK:块级迁移 MIGRATE_FILE:文件级迁移
	Type *TaskByServerSourceType `json:"type,omitempty"`

	// 任务状态
	State *string `json:"state,omitempty"`

	// 开始时间
	StartDate *int64 `json:"start_date,omitempty"`

	// 限速
	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	// 迁移速率
	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	// 是否启动虚拟机
	StartTargetServer *bool `json:"start_target_server,omitempty"`

	// 虚拟机模板ID
	VmTemplateId *string `json:"vm_template_id,omitempty"`

	// region_id
	RegionId *string `json:"region_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	TargetServer *TargetServerById `json:"target_server,omitempty"`

	// 日志收集状态 INIT:就绪 UPLOADING:上传中 UPLOAD_FAIL:上传失败 UPLOADED:已上传
	LogCollectStatus *TaskByServerSourceLogCollectStatus `json:"log_collect_status,omitempty"`

	// 是否使用已有虚拟机
	ExistServer *bool `json:"exist_server,omitempty"`

	// 是否使用公网IP
	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	CloneServer *CloneServer `json:"clone_server,omitempty"`

	// 当前子任务及进度
	SubtaskInfo *string `json:"subtask_info,omitempty"`
}

func (o TaskByServerSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskByServerSource struct{}"
	}

	return strings.Join([]string{"TaskByServerSource", string(data)}, " ")
}

type TaskByServerSourceType struct {
	value string
}

type TaskByServerSourceTypeEnum struct {
	MIGRATE_BLOCK TaskByServerSourceType
	MIGRATE_FILE  TaskByServerSourceType
}

func GetTaskByServerSourceTypeEnum() TaskByServerSourceTypeEnum {
	return TaskByServerSourceTypeEnum{
		MIGRATE_BLOCK: TaskByServerSourceType{
			value: "MIGRATE_BLOCK",
		},
		MIGRATE_FILE: TaskByServerSourceType{
			value: "MIGRATE_FILE",
		},
	}
}

func (c TaskByServerSourceType) Value() string {
	return c.value
}

func (c TaskByServerSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskByServerSourceType) UnmarshalJSON(b []byte) error {
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

type TaskByServerSourceLogCollectStatus struct {
	value string
}

type TaskByServerSourceLogCollectStatusEnum struct {
	INIT        TaskByServerSourceLogCollectStatus
	UPLOADING   TaskByServerSourceLogCollectStatus
	UPLOAD_FAIL TaskByServerSourceLogCollectStatus
	UPLOADED    TaskByServerSourceLogCollectStatus
}

func GetTaskByServerSourceLogCollectStatusEnum() TaskByServerSourceLogCollectStatusEnum {
	return TaskByServerSourceLogCollectStatusEnum{
		INIT: TaskByServerSourceLogCollectStatus{
			value: "INIT",
		},
		UPLOADING: TaskByServerSourceLogCollectStatus{
			value: "UPLOADING",
		},
		UPLOAD_FAIL: TaskByServerSourceLogCollectStatus{
			value: "UPLOAD_FAIL",
		},
		UPLOADED: TaskByServerSourceLogCollectStatus{
			value: "UPLOADED",
		},
	}
}

func (c TaskByServerSourceLogCollectStatus) Value() string {
	return c.value
}

func (c TaskByServerSourceLogCollectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskByServerSourceLogCollectStatus) UnmarshalJSON(b []byte) error {
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
