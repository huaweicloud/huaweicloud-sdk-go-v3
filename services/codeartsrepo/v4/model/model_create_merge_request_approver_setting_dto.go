package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMergeRequestApproverSettingDto 审核策配置
type CreateMergeRequestApproverSettingDto struct {

	// 配置分支，可使用*作为通配符使用，如：dev* 指以dev开头的所有分支
	Target *string `json:"target,omitempty"`

	// 为false时，“最小检视人数”，“最小审核人数”，“重新推送代码时重置审核人”，“重新推送代码时重置检视人”，“仅能从以下审核人/检视人中追加审核人/检视人”置为默认
	IsUseApproval *bool `json:"is_use_approval,omitempty"`

	// 最小审核人数
	ApprovalRequiredApprovers *int32 `json:"approval_required_approvers,omitempty"`

	// 最小检视人数
	ApprovalRequiredReviewers *int32 `json:"approval_required_reviewers,omitempty"`

	// 推送时是否重置审核门禁状态
	ResetApprovalsOnPush *bool `json:"reset_approvals_on_push,omitempty"`

	// 推送时是否重置检视门禁状态
	ResetReviewersOnPush *bool `json:"reset_reviewers_on_push,omitempty"`

	// 是否开启仅能从以下审核/检视人中追加审核人/检视人
	ApproversFromProject *bool `json:"approvers_from_project,omitempty"`

	// 追加检视人人用户id列表
	AppendReviewerIds *[]int32 `json:"append_reviewer_ids,omitempty"`

	// 追加审核人用户id列表
	AppendApproverIds *[]int32 `json:"append_approver_ids,omitempty"`

	// 是否开启流水线门禁
	OnlyMergeWhenPipelinePass *bool `json:"only_merge_when_pipeline_pass,omitempty"`

	// 合并人用户id列表
	AssigneeIds *[]int32 `json:"assignee_ids,omitempty"`

	// 审核人用户id列表
	ApproverIds *[]int32 `json:"approver_ids,omitempty"`

	// 检视人用户id列表
	ReviewerIds *[]int32 `json:"reviewer_ids,omitempty"`
}

func (o CreateMergeRequestApproverSettingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestApproverSettingDto struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestApproverSettingDto", string(data)}, " ")
}
