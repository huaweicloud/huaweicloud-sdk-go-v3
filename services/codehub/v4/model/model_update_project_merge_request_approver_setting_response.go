package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateProjectMergeRequestApproverSettingResponse Response Object
type UpdateProjectMergeRequestApproverSettingResponse struct {

	// **参数解释：**  设置主键id。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  分支，可使用*作为通配符使用，如：dev* 指以dev开头的所有分支
	Target *string `json:"target,omitempty"`

	// 设置类型，当前仅开放branch类型
	TargetType *UpdateProjectMergeRequestApproverSettingResponseTargetType `json:"target_type,omitempty"`

	// 是否为审核模式，审核模式：true 评分模式：false
	IsUseApproval *bool `json:"is_use_approval,omitempty"`

	// 最小检视人数
	ApprovalRequiredReviewers *int32 `json:"approval_required_reviewers,omitempty"`

	// 最小审核人数
	ApprovalRequiredApprovers *int32 `json:"approval_required_approvers,omitempty"`

	// 推送时是否重置审核门禁状态
	ResetApprovalsOnPush *bool `json:"reset_approvals_on_push,omitempty"`

	// 推送时是否重置检视门禁状态
	ResetReviewersOnPush *bool `json:"reset_reviewers_on_push,omitempty"`

	// 是否开启仅能从以下审核/检视人中追加审核人/检视人
	ApproversFromProject *bool `json:"approvers_from_project,omitempty"`

	// 追加检视人id列表
	AppendReviewerIds *[]int32 `json:"append_reviewer_ids,omitempty"`

	// 追加检视人实体列表
	AppendReviewers *[]UserBasicDto `json:"append_reviewers,omitempty"`

	// 追加审核人id列表
	AppendApproverIds *[]int32 `json:"append_approver_ids,omitempty"`

	// 追加审核人实体列表
	AppendApprovers *[]UserBasicDto `json:"append_approvers,omitempty"`

	// 是否开启流水线门禁
	OnlyMergeWhenPipelinePass *bool `json:"only_merge_when_pipeline_pass,omitempty"`

	// 合并人id列表
	AssigneeIds *[]int32 `json:"assignee_ids,omitempty"`

	// 合并人实体列表
	Assignees *[]UserBasicDto `json:"assignees,omitempty"`

	// 审核人id列表
	ApproverIds *[]int32 `json:"approver_ids,omitempty"`

	// 审核人实体列表
	Approvers *[]UserBasicDto `json:"approvers,omitempty"`

	// 检视人id列表
	ReviewerIds *[]int32 `json:"reviewer_ids,omitempty"`

	// 检视人实体列表
	Reviewers      *[]UserBasicDto `json:"reviewers,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateProjectMergeRequestApproverSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectMergeRequestApproverSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectMergeRequestApproverSettingResponse", string(data)}, " ")
}

type UpdateProjectMergeRequestApproverSettingResponseTargetType struct {
	value string
}

type UpdateProjectMergeRequestApproverSettingResponseTargetTypeEnum struct {
	BRANCH UpdateProjectMergeRequestApproverSettingResponseTargetType
	FILE   UpdateProjectMergeRequestApproverSettingResponseTargetType
	BINARY UpdateProjectMergeRequestApproverSettingResponseTargetType
}

func GetUpdateProjectMergeRequestApproverSettingResponseTargetTypeEnum() UpdateProjectMergeRequestApproverSettingResponseTargetTypeEnum {
	return UpdateProjectMergeRequestApproverSettingResponseTargetTypeEnum{
		BRANCH: UpdateProjectMergeRequestApproverSettingResponseTargetType{
			value: "branch",
		},
		FILE: UpdateProjectMergeRequestApproverSettingResponseTargetType{
			value: "file",
		},
		BINARY: UpdateProjectMergeRequestApproverSettingResponseTargetType{
			value: "binary",
		},
	}
}

func (c UpdateProjectMergeRequestApproverSettingResponseTargetType) Value() string {
	return c.value
}

func (c UpdateProjectMergeRequestApproverSettingResponseTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProjectMergeRequestApproverSettingResponseTargetType) UnmarshalJSON(b []byte) error {
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
