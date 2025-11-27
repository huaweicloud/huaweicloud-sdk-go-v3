package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskByServerSources 源端列表中关联的任务
type TaskByServerSources struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 迁移项目类型 MIGRATE_BLOCK:块级迁移 MIGRATE_FILE:文件级迁移
	Type *TaskByServerSourcesType `json:"type,omitempty"`

	// 任务状态
	State *string `json:"state,omitempty"`

	// 预估结束时间
	EstimateCompleteTime *int64 `json:"estimate_complete_time,omitempty"`

	// 开始时间
	StartDate *int64 `json:"start_date,omitempty"`

	// 限速
	SpeedLimit *int32 `json:"speed_limit,omitempty"`

	// 迁移速率
	MigrateSpeed *float64 `json:"migrate_speed,omitempty"`

	// 压缩率
	CompressRate *float64 `json:"compress_rate,omitempty"`

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
	LogCollectStatus *TaskByServerSourcesLogCollectStatus `json:"log_collect_status,omitempty"`

	// 是否使用已有虚拟机
	ExistServer *bool `json:"exist_server,omitempty"`

	// 是否使用公网IP
	UsePublicIp *bool `json:"use_public_ip,omitempty"`

	CloneServer *CloneServer `json:"clone_server,omitempty"`

	// 迁移剩余时间（秒）
	RemainSeconds *int64 `json:"remain_seconds,omitempty"`

	// 上传日志指定桶名称
	LogBucket *string `json:"log_bucket,omitempty"`

	// 分享链接有效期
	LogExpire *int64 `json:"log_expire,omitempty"`

	// 日志上传时间
	LogUploadTime *int64 `json:"log_upload_time,omitempty"`

	// 分享链接url
	LogShareUrl *string `json:"log_share_url,omitempty"`

	// 当前子任务及进度
	SubtaskInfo *string `json:"subtask_info,omitempty"`
}

func (o TaskByServerSources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskByServerSources struct{}"
	}

	return strings.Join([]string{"TaskByServerSources", string(data)}, " ")
}

type TaskByServerSourcesType struct {
	value string
}

type TaskByServerSourcesTypeEnum struct {
	MIGRATE_BLOCK TaskByServerSourcesType
	MIGRATE_FILE  TaskByServerSourcesType
}

func GetTaskByServerSourcesTypeEnum() TaskByServerSourcesTypeEnum {
	return TaskByServerSourcesTypeEnum{
		MIGRATE_BLOCK: TaskByServerSourcesType{
			value: "MIGRATE_BLOCK",
		},
		MIGRATE_FILE: TaskByServerSourcesType{
			value: "MIGRATE_FILE",
		},
	}
}

func (c TaskByServerSourcesType) Value() string {
	return c.value
}

func (c TaskByServerSourcesType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskByServerSourcesType) UnmarshalJSON(b []byte) error {
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

type TaskByServerSourcesLogCollectStatus struct {
	value string
}

type TaskByServerSourcesLogCollectStatusEnum struct {
	INIT        TaskByServerSourcesLogCollectStatus
	UPLOADING   TaskByServerSourcesLogCollectStatus
	UPLOAD_FAIL TaskByServerSourcesLogCollectStatus
	UPLOADED    TaskByServerSourcesLogCollectStatus
}

func GetTaskByServerSourcesLogCollectStatusEnum() TaskByServerSourcesLogCollectStatusEnum {
	return TaskByServerSourcesLogCollectStatusEnum{
		INIT: TaskByServerSourcesLogCollectStatus{
			value: "INIT",
		},
		UPLOADING: TaskByServerSourcesLogCollectStatus{
			value: "UPLOADING",
		},
		UPLOAD_FAIL: TaskByServerSourcesLogCollectStatus{
			value: "UPLOAD_FAIL",
		},
		UPLOADED: TaskByServerSourcesLogCollectStatus{
			value: "UPLOADED",
		},
	}
}

func (c TaskByServerSourcesLogCollectStatus) Value() string {
	return c.value
}

func (c TaskByServerSourcesLogCollectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskByServerSourcesLogCollectStatus) UnmarshalJSON(b []byte) error {
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
