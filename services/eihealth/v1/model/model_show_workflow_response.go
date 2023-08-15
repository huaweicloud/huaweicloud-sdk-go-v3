package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowResponse Response Object
type ShowWorkflowResponse struct {

	// 流程id
	Id *string `json:"id,omitempty"`

	// 流程名称
	Name *string `json:"name,omitempty"`

	// 流程版本
	Version *string `json:"version,omitempty"`

	// 流程简述
	Summary *string `json:"summary,omitempty"`

	// 流程描述
	Description *string `json:"description,omitempty"`

	// 流程标签
	Labels *[]string `json:"labels,omitempty"`

	// 超时时间
	Timeout *int32 `json:"timeout,omitempty"`

	// 流程的输出路径
	OutputDir *string `json:"output_dir,omitempty"`

	// 流程的子任务
	Tasks *[]TaskDefinitionDto `json:"tasks,omitempty"`

	// 流程创建时刻的应用快照，当query填workflow_snapshot_sign有值;K为appId,V为sign
	AppSnapshotSign map[string]string `json:"app_snapshot_sign,omitempty"`

	// 流程的创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 流程的更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源资源id
	SourceResourceId *string `json:"source_resource_id,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowResponse", string(data)}, " ")
}
