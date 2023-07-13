package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NextflowWorkflowListDto struct {

	// 流程id
	Id *string `json:"id,omitempty"`

	// 流程名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 流程标签
	Labels *[]string `json:"labels,omitempty"`

	// 创建流程时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新流程时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 创建用户名称
	Creator *string `json:"creator,omitempty"`

	// 源项目名称
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 源资源id
	SourceResourceId *string `json:"source_resource_id,omitempty"`
}

func (o NextflowWorkflowListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextflowWorkflowListDto struct{}"
	}

	return strings.Join([]string{"NextflowWorkflowListDto", string(data)}, " ")
}
