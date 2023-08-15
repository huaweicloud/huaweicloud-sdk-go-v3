package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobRequest Request Object
type ListJobRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 最大开始时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 作业名称 取值范围：长度为[1,63]，以小写字母开头，允许出现中划线(-)、小写字母和数字，且必须以小写字母或数字结尾。
	JobName *string `json:"job_name,omitempty"`

	// 标签列表
	Labels *[]string `json:"labels,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`

	// 排序规则 目前默认时间降序
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序规则 目前默认时间降序，支持根据status
	SortKey *string `json:"sort_key,omitempty"`

	// 最小开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 作业运行状态 取值（Succeeded|Running|Pending|Failed|Cancelling|Cancelled|Unknown）
	Status *string `json:"status,omitempty"`

	// 作业依赖的组件名称(有可能是Workflow，有可能是app), 取值范围：长度为[1,56]，以小写字母开头，允许出现中划线(-)、小写字母和数字，且必须以小写字母或数字结尾。
	ToolName *string `json:"tool_name,omitempty"`

	// 作业创建者
	UserName *string `json:"user_name,omitempty"`

	// 最小结束时间
	FinishStartTime *int64 `json:"finish_start_time,omitempty"`

	// 最大结束时间
	FinishEndTime *int64 `json:"finish_end_time,omitempty"`
}

func (o ListJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobRequest struct{}"
	}

	return strings.Join([]string{"ListJobRequest", string(data)}, " ")
}
