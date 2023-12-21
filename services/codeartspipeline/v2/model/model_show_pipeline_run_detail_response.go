package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineRunDetailResponse Response Object
type ShowPipelineRunDetailResponse struct {

	// 流水线运行实例ID
	Id *string `json:"id,omitempty"`

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 流水线版本
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// 流水线名称
	Name *string `json:"name,omitempty"`

	// 运行描述
	Description *string `json:"description,omitempty"`

	// 是否为变更流水线
	IsPublish *bool `json:"is_publish,omitempty"`

	// 运行人ID
	ExecutorId *string `json:"executor_id,omitempty"`

	// 运行人名称
	ExecutorName *string `json:"executor_name,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	// 运行序号
	RunNumber *int32 `json:"run_number,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 阶段信息
	Stages *[]StageRun `json:"stages,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 局点
	Region *string `json:"region,omitempty"`

	// 组件ID
	ComponentId *string `json:"component_id,omitempty"`

	// 语言
	Language *string `json:"language,omitempty"`

	// 运行源信息
	Sources *[]RunPipelineSource `json:"sources,omitempty"`

	// 流水线运行产物
	Artifacts *[]PackageInfo `json:"artifacts,omitempty"`

	// 流水线运行实例ID
	SubjectId *string `json:"subject_id,omitempty"`

	// 分组ID
	GroupId *string `json:"group_id,omitempty"`

	// 分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 详情页地址
	DetailUrl *string `json:"detail_url,omitempty"`

	// 当前系统时间
	CurrentSystemTime *int64 `json:"current_system_time,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowPipelineRunDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineRunDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineRunDetailResponse", string(data)}, " ")
}
