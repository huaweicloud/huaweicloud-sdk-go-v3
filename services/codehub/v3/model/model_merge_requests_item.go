package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeRequestsItem struct {
	Author *Author `json:"author,omitempty"`

	// 关闭时间
	ClosedAt *string `json:"closed_at,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 合并请求描述
	Description *string `json:"description,omitempty"`

	// 源分支
	DevcloudSourceBranch *string `json:"devcloud_source_branch,omitempty"`

	// 合并请求id
	Id *float64 `json:"id,omitempty"`

	// 当前仓库内合并请求序号
	Iid *float64 `json:"iid,omitempty"`

	// 检视人
	MergeRequestAssigneeList *[]Author `json:"merge_request_assignee_list,omitempty"`

	// 是否可以被合并
	MergeStatus *string `json:"merge_status,omitempty"`

	// 源分支
	SourceBranch *string `json:"source_branch,omitempty"`

	// 合并请求状态
	State *string `json:"state,omitempty"`

	// 目标分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o MergeRequestsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestsItem struct{}"
	}

	return strings.Join([]string{"MergeRequestsItem", string(data)}, " ")
}
