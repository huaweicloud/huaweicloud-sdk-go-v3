package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserWorkflowRequest Request Object
type ListUserWorkflowRequest struct {

	// 是否仅展示本人创建资源。
	IsCreator *bool `json:"is_creator,omitempty"`

	// 流程ID，仅支持精确搜索。
	Id *string `json:"id,omitempty"`

	// 流程名称，支持模糊搜索。
	Name *string `json:"name,omitempty"`

	// 短描述，支持模糊搜索。
	Summary *string `json:"summary,omitempty"`

	// 空间名称列表，多个值之间使用逗号分隔。
	EihealthProjectNames *[]string `json:"eihealth_project_names,omitempty"`

	// 标签列表。
	Labels *[]string `json:"labels,omitempty"`

	// 开始创建时间。
	StartCreateTime *int64 `json:"start_create_time,omitempty"`

	// 结束创建时间。
	EndCreateTime *int64 `json:"end_create_time,omitempty"`

	// 开始更新时间。
	StartUpdateTime *int64 `json:"start_update_time,omitempty"`

	// 结束更新时间。
	EndUpdateTime *int64 `json:"end_update_time,omitempty"`

	// 排序规则。
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序规则，默认根据创建时间降序，支持根据create_time|update_time排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListUserWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ListUserWorkflowRequest", string(data)}, " ")
}
