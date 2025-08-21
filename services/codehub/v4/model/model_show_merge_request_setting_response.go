package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowMergeRequestSettingResponse Response Object
type ShowMergeRequestSettingResponse struct {

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
	ReviewMode *ShowMergeRequestSettingResponseReviewMode `json:"review_mode,omitempty"`

	// **参数解释：** 允许合并请求合并或关闭后继续做代码检视和评论
	AddNotesAfterMerged *bool `json:"add_notes_after_merged,omitempty"`

	// **参数解释：** merged_by：使用MR合入者生成Merge Commit created_by：使用MR创建者生成Merge Commit
	MergedCommitAuthor *ShowMergeRequestSettingResponseMergedCommitAuthor `json:"merged_commit_author,omitempty"`

	// **参数解释：** 是否将星级评价作为合入门禁-参与星级评价角色。 29：勾选开发者 30：勾选开发者，Committer 35：勾选Committer 40：都不勾选
	EvaluationRole *ShowMergeRequestSettingResponseEvaluationRole `json:"evaluation_role,omitempty"`

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
	MergeMethod *ShowMergeRequestSettingResponseMergeMethod `json:"merge_method,omitempty"`

	// **参数解释：** 打分模式最低合入分数。
	OnlyAllowMergeIfVoteBiggerThan *int32 `json:"only_allow_merge_if_vote_bigger_than,omitempty"`

	// **参数解释：** 仅合并人和合并合并请求。
	OnlyAssigneeCanMerge *bool `json:"only_assignee_can_merge,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId   *int32 `json:"repository_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMergeRequestSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeRequestSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowMergeRequestSettingResponse", string(data)}, " ")
}

type ShowMergeRequestSettingResponseReviewMode struct {
	value string
}

type ShowMergeRequestSettingResponseReviewModeEnum struct {
	APPROVAL ShowMergeRequestSettingResponseReviewMode
	VOTE     ShowMergeRequestSettingResponseReviewMode
}

func GetShowMergeRequestSettingResponseReviewModeEnum() ShowMergeRequestSettingResponseReviewModeEnum {
	return ShowMergeRequestSettingResponseReviewModeEnum{
		APPROVAL: ShowMergeRequestSettingResponseReviewMode{
			value: "approval",
		},
		VOTE: ShowMergeRequestSettingResponseReviewMode{
			value: "vote",
		},
	}
}

func (c ShowMergeRequestSettingResponseReviewMode) Value() string {
	return c.value
}

func (c ShowMergeRequestSettingResponseReviewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestSettingResponseReviewMode) UnmarshalJSON(b []byte) error {
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

type ShowMergeRequestSettingResponseMergedCommitAuthor struct {
	value string
}

type ShowMergeRequestSettingResponseMergedCommitAuthorEnum struct {
	MERGED_BY  ShowMergeRequestSettingResponseMergedCommitAuthor
	CREATED_BY ShowMergeRequestSettingResponseMergedCommitAuthor
}

func GetShowMergeRequestSettingResponseMergedCommitAuthorEnum() ShowMergeRequestSettingResponseMergedCommitAuthorEnum {
	return ShowMergeRequestSettingResponseMergedCommitAuthorEnum{
		MERGED_BY: ShowMergeRequestSettingResponseMergedCommitAuthor{
			value: "merged_by",
		},
		CREATED_BY: ShowMergeRequestSettingResponseMergedCommitAuthor{
			value: "created_by",
		},
	}
}

func (c ShowMergeRequestSettingResponseMergedCommitAuthor) Value() string {
	return c.value
}

func (c ShowMergeRequestSettingResponseMergedCommitAuthor) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestSettingResponseMergedCommitAuthor) UnmarshalJSON(b []byte) error {
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

type ShowMergeRequestSettingResponseEvaluationRole struct {
	value int32
}

type ShowMergeRequestSettingResponseEvaluationRoleEnum struct {
	E_29 ShowMergeRequestSettingResponseEvaluationRole
	E_30 ShowMergeRequestSettingResponseEvaluationRole
	E_35 ShowMergeRequestSettingResponseEvaluationRole
	E_40 ShowMergeRequestSettingResponseEvaluationRole
}

func GetShowMergeRequestSettingResponseEvaluationRoleEnum() ShowMergeRequestSettingResponseEvaluationRoleEnum {
	return ShowMergeRequestSettingResponseEvaluationRoleEnum{
		E_29: ShowMergeRequestSettingResponseEvaluationRole{
			value: 29,
		}, E_30: ShowMergeRequestSettingResponseEvaluationRole{
			value: 30,
		}, E_35: ShowMergeRequestSettingResponseEvaluationRole{
			value: 35,
		}, E_40: ShowMergeRequestSettingResponseEvaluationRole{
			value: 40,
		},
	}
}

func (c ShowMergeRequestSettingResponseEvaluationRole) Value() int32 {
	return c.value
}

func (c ShowMergeRequestSettingResponseEvaluationRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestSettingResponseEvaluationRole) UnmarshalJSON(b []byte) error {
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

type ShowMergeRequestSettingResponseMergeMethod struct {
	value string
}

type ShowMergeRequestSettingResponseMergeMethodEnum struct {
	MERGE        ShowMergeRequestSettingResponseMergeMethod
	REBASE_MERGE ShowMergeRequestSettingResponseMergeMethod
	FF           ShowMergeRequestSettingResponseMergeMethod
}

func GetShowMergeRequestSettingResponseMergeMethodEnum() ShowMergeRequestSettingResponseMergeMethodEnum {
	return ShowMergeRequestSettingResponseMergeMethodEnum{
		MERGE: ShowMergeRequestSettingResponseMergeMethod{
			value: "merge",
		},
		REBASE_MERGE: ShowMergeRequestSettingResponseMergeMethod{
			value: "rebase_merge",
		},
		FF: ShowMergeRequestSettingResponseMergeMethod{
			value: "ff",
		},
	}
}

func (c ShowMergeRequestSettingResponseMergeMethod) Value() string {
	return c.value
}

func (c ShowMergeRequestSettingResponseMergeMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMergeRequestSettingResponseMergeMethod) UnmarshalJSON(b []byte) error {
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
