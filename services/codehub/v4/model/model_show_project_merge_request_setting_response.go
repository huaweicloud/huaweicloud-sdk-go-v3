package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowProjectMergeRequestSettingResponse Response Object
type ShowProjectMergeRequestSettingResponse struct {

	// **参数解释：** 主键id
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 禁止合入自己创建的合并请求
	DisableMergeBySelf *bool `json:"disable_merge_by_self,omitempty"`

	// **参数解释：** 禁止审核自己创建的合并请求
	DisableApproveBySelf *bool `json:"disable_approve_by_self,omitempty"`

	// **参数解释：** 禁止检视自己创建的合并请求
	DisableReviewBySelf *bool `json:"disable_review_by_self,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 允许仓库管理员及项目经理强制合入
	CanForceMerge *bool `json:"can_force_merge,omitempty"`

	// **参数解释：** 禁止Squash合并（合入MR时禁止Squash合并）
	DisableSquashMerge *bool `json:"disable_squash_merge,omitempty"`

	// **参数解释：** 必须与CodeArts Req关联
	MustRelateIssue *bool `json:"must_relate_issue,omitempty"`

	// **参数解释：** 必须与CodeArts Req关联-选择目标分支配置合并请求策略(分支以逗号分支的字符串)
	NeedRelateIssueBranches *string `json:"need_relate_issue_branches,omitempty"`

	// **参数解释：** 必须与CodeArts Req关联-所有E2E单号校验必须通过
	NeedAllIssuesCheckPassed *bool `json:"need_all_issues_check_passed,omitempty"`

	// **参数解释：** 合入模式
	ReviewMode *ShowProjectMergeRequestSettingResponseReviewMode `json:"review_mode,omitempty"`

	// **参数解释：** 允许合并请求合并或关闭后继续做代码检视和评论
	AddNotesAfterMerged *bool `json:"add_notes_after_merged,omitempty"`

	// **参数解释：** merged_by：使用MR合入者生成Merge Commit created_by：使用MR创建者生成Merge Commit
	MergedCommitAuthor *ShowProjectMergeRequestSettingResponseMergedCommitAuthor `json:"merged_commit_author,omitempty"`

	// **参数解释：** 是否将星级评价作为合入门禁-参与星级评价角色。 29：勾选开发者 30：勾选开发者，Committer 35：勾选Committer 40：都不勾选
	EvaluationRole *ShowProjectMergeRequestSettingResponseEvaluationRole `json:"evaluation_role,omitempty"`

	// **参数解释：** 是否将星级评价作为合入门禁
	EvaluationMergeGate *bool `json:"evaluation_merge_gate,omitempty"`

	// **参数解释：** 是否将自动合并的MR状态标记为关闭状态
	MarkAutoMergedMrAsClosed *bool `json:"mark_auto_merged_mr_as_closed,omitempty"`

	// **参数解释：** 新建合并请求，默认开启合并后删除源分支
	DeleteSourceBranchWhenMerged *bool `json:"delete_source_branch_when_merged,omitempty"`

	// **参数解释：** 新建合并请求，默认开启Squash合并
	AutoSquashMerge *bool `json:"auto_squash_merge,omitempty"`

	// **参数解释：** Squash合并不产生Merge节点
	SquashMergeWithNoMergeCommit *bool `json:"squash_merge_with_no_merge_commit,omitempty"`

	// **参数解释：** 只能关联一个单号
	OnlyAllowOneIssueCheckPassed *bool `json:"only_allow_one_issue_check_passed,omitempty"`

	// **参数解释：** 启用MR自定义评价维度分类（MR评价设置）
	EnableCustomEvaluation *bool `json:"enable_custom_evaluation,omitempty"`

	// **参数解释：** 评价维度（MR评价设置）
	EvaluationTypes *[]EvaluationTypeDto `json:"evaluation_types,omitempty"`

	// **参数解释：** 不能重新打开一个已经关闭的合并请求
	CanReopen *bool `json:"can_reopen,omitempty"`

	// **参数解释：** 评审问题全部解决才能合入
	OnlyAllowMergeIfAllDiscussionsAreResolved *bool `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`

	// **参数解释：** 合并模式。 merge：通过 Merge Commit 合并 rebase_merge：通过 Merge Commit 合并(记录半线性历史) ff：Fast-forward 合并
	MergeMethod *ShowProjectMergeRequestSettingResponseMergeMethod `json:"merge_method,omitempty"`

	// **参数解释：** 打分模式最低合入分数。
	OnlyAllowMergeIfVoteBiggerThan *int32 `json:"only_allow_merge_if_vote_bigger_than,omitempty"`

	// **参数解释：** 仅合并人和合并合并请求。
	OnlyAssigneeCanMerge *bool `json:"only_assignee_can_merge,omitempty"`

	// **参数解释：** 项目id。
	ProjectId      *string `json:"project_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProjectMergeRequestSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectMergeRequestSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectMergeRequestSettingResponse", string(data)}, " ")
}

type ShowProjectMergeRequestSettingResponseReviewMode struct {
	value string
}

type ShowProjectMergeRequestSettingResponseReviewModeEnum struct {
	APPROVAL ShowProjectMergeRequestSettingResponseReviewMode
	VOTE     ShowProjectMergeRequestSettingResponseReviewMode
}

func GetShowProjectMergeRequestSettingResponseReviewModeEnum() ShowProjectMergeRequestSettingResponseReviewModeEnum {
	return ShowProjectMergeRequestSettingResponseReviewModeEnum{
		APPROVAL: ShowProjectMergeRequestSettingResponseReviewMode{
			value: "approval",
		},
		VOTE: ShowProjectMergeRequestSettingResponseReviewMode{
			value: "vote",
		},
	}
}

func (c ShowProjectMergeRequestSettingResponseReviewMode) Value() string {
	return c.value
}

func (c ShowProjectMergeRequestSettingResponseReviewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectMergeRequestSettingResponseReviewMode) UnmarshalJSON(b []byte) error {
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

type ShowProjectMergeRequestSettingResponseMergedCommitAuthor struct {
	value string
}

type ShowProjectMergeRequestSettingResponseMergedCommitAuthorEnum struct {
	MERGED_BY  ShowProjectMergeRequestSettingResponseMergedCommitAuthor
	CREATED_BY ShowProjectMergeRequestSettingResponseMergedCommitAuthor
}

func GetShowProjectMergeRequestSettingResponseMergedCommitAuthorEnum() ShowProjectMergeRequestSettingResponseMergedCommitAuthorEnum {
	return ShowProjectMergeRequestSettingResponseMergedCommitAuthorEnum{
		MERGED_BY: ShowProjectMergeRequestSettingResponseMergedCommitAuthor{
			value: "merged_by",
		},
		CREATED_BY: ShowProjectMergeRequestSettingResponseMergedCommitAuthor{
			value: "created_by",
		},
	}
}

func (c ShowProjectMergeRequestSettingResponseMergedCommitAuthor) Value() string {
	return c.value
}

func (c ShowProjectMergeRequestSettingResponseMergedCommitAuthor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectMergeRequestSettingResponseMergedCommitAuthor) UnmarshalJSON(b []byte) error {
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

type ShowProjectMergeRequestSettingResponseEvaluationRole struct {
	value int32
}

type ShowProjectMergeRequestSettingResponseEvaluationRoleEnum struct {
	E_29 ShowProjectMergeRequestSettingResponseEvaluationRole
	E_30 ShowProjectMergeRequestSettingResponseEvaluationRole
	E_35 ShowProjectMergeRequestSettingResponseEvaluationRole
	E_40 ShowProjectMergeRequestSettingResponseEvaluationRole
}

func GetShowProjectMergeRequestSettingResponseEvaluationRoleEnum() ShowProjectMergeRequestSettingResponseEvaluationRoleEnum {
	return ShowProjectMergeRequestSettingResponseEvaluationRoleEnum{
		E_29: ShowProjectMergeRequestSettingResponseEvaluationRole{
			value: 29,
		}, E_30: ShowProjectMergeRequestSettingResponseEvaluationRole{
			value: 30,
		}, E_35: ShowProjectMergeRequestSettingResponseEvaluationRole{
			value: 35,
		}, E_40: ShowProjectMergeRequestSettingResponseEvaluationRole{
			value: 40,
		},
	}
}

func (c ShowProjectMergeRequestSettingResponseEvaluationRole) Value() int32 {
	return c.value
}

func (c ShowProjectMergeRequestSettingResponseEvaluationRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectMergeRequestSettingResponseEvaluationRole) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ShowProjectMergeRequestSettingResponseMergeMethod struct {
	value string
}

type ShowProjectMergeRequestSettingResponseMergeMethodEnum struct {
	MERGE        ShowProjectMergeRequestSettingResponseMergeMethod
	REBASE_MERGE ShowProjectMergeRequestSettingResponseMergeMethod
	FF           ShowProjectMergeRequestSettingResponseMergeMethod
}

func GetShowProjectMergeRequestSettingResponseMergeMethodEnum() ShowProjectMergeRequestSettingResponseMergeMethodEnum {
	return ShowProjectMergeRequestSettingResponseMergeMethodEnum{
		MERGE: ShowProjectMergeRequestSettingResponseMergeMethod{
			value: "merge",
		},
		REBASE_MERGE: ShowProjectMergeRequestSettingResponseMergeMethod{
			value: "rebase_merge",
		},
		FF: ShowProjectMergeRequestSettingResponseMergeMethod{
			value: "ff",
		},
	}
}

func (c ShowProjectMergeRequestSettingResponseMergeMethod) Value() string {
	return c.value
}

func (c ShowProjectMergeRequestSettingResponseMergeMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectMergeRequestSettingResponseMergeMethod) UnmarshalJSON(b []byte) error {
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
