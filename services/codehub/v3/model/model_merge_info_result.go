package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeInfoResult struct {

	// 评审人
	ApprovalMergeRequestApprovers *[]ApprovalMergeRequestApproversItem `json:"approvalMergeRequestApprovers,omitempty"`

	Author *Author `json:"author,omitempty"`

	// 关闭时间
	ClosedAt *string `json:"closedAt,omitempty"`

	// 创建时间
	CreatedAt *string `json:"createdAt,omitempty"`

	// 合并请求描述
	Description *string `json:"description,omitempty"`

	// 源分支
	DevcloudSourceBranch *string `json:"devcloudSourceBranch,omitempty"`

	// 合并请求id
	Id *float64 `json:"id,omitempty"`

	// 当前仓库内合并请求序号
	Iid *float64 `json:"iid,omitempty"`

	// 源分支是否存在
	IsSourceBranchExist *bool `json:"isSourceBranchExist,omitempty"`

	// 检视人
	MergeRequestAssigneeList *[]MergeRequestAssigneeListItem `json:"mergeRequestAssigneeList,omitempty"`

	MergeRequestDiff *MergeRequestDiff `json:"mergeRequestDiff,omitempty"`

	// 是否可以被合并
	MergeStatus *string `json:"mergeStatus,omitempty"`

	// 源分支
	SourceBranch *string `json:"sourceBranch,omitempty"`

	// 合并请求状态
	State *string `json:"state,omitempty"`

	// 目标分支
	TargetBranch *string `json:"targetBranch,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (o MergeInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeInfoResult struct{}"
	}

	return strings.Join([]string{"MergeInfoResult", string(data)}, " ")
}
