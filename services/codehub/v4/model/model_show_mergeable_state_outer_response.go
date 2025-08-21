package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMergeableStateOuterResponse Response Object
type ShowMergeableStateOuterResponse struct {

	// **参数解释：** 合并请求id。
	MergeRequestId *int32 `json:"merge_request_id,omitempty"`

	// **参数解释：** 合并请求的可合入状态。 **约束限制：** - true，可合入。 - false，不可合入。
	State *bool `json:"state,omitempty"`

	// **参数解释：** 合并请求冲突门禁是否通过。 **约束限制：** - true，不存在冲突。 - false，存在冲突。
	ConflictPassed *bool `json:"conflict_passed,omitempty"`

	// **参数解释：** 合并请求是否需要变基。 **约束限制：** - true，不需变基。 - false，需要变基。
	NonFfPassed *bool `json:"non_ff_passed,omitempty"`

	// **参数解释：** 对当前用户是否有合入当前合并请求的权限判断。 **约束限制：** - true，有权限合入。 - false，无权限合入。
	MergedByUserPassed *bool `json:"merged_by_user_passed,omitempty"`

	// **参数解释：** 合并请求WIP门禁是否通过。 **约束限制：** - true，非WIP状态。 - false，WIP状态。
	WorkInProgressPassed *bool `json:"work_in_progress_passed,omitempty"`

	// **参数解释：** 合并请求检视意见门禁是否通过。 **约束限制：** - true，检视意见门禁通过。 - false，检视意见门禁不通过。
	ResolveDiscussionPassed *bool `json:"resolve_discussion_passed,omitempty"`

	// **参数解释：** 合并请求流水线门禁是否通过。 **约束限制：** - true，合并请求流水线门禁通过。 - false，合并请求流水线门禁不通过。
	CiStatePassed *bool `json:"ci_state_passed,omitempty"`

	// **参数解释：** 对当前用户是否有合入自己创建的合并请求的判断。 **约束限制：** - true，非自己创建的mr可以合入。 - false，自己创建的mr不能合入。
	MergeBySelfPassed *bool `json:"merge_by_self_passed,omitempty"`

	// **参数解释：** 当前用户是否可以强制合入当前合并请求。 **约束限制：** - true，可以强制合入。 - false，不能强制合入。
	CanForceMerge *bool `json:"can_force_merge,omitempty"`

	// **参数解释：** 合并请求评分门禁是否通过。 **约束限制：** - true，评分门禁通过。 - false，评分门禁不通过。
	VotePassed *bool `json:"vote_passed,omitempty"`

	// **参数解释：** 合并请求必须与CodeArts Req关联门禁是否通过。 **约束限制：** - true，合并请求必须与CodeArts Req关联门禁通过。 - false，合并请求必须与CodeArts Req关联门禁不通过。
	E2eCheckPassed *bool `json:"e2e_check_passed,omitempty"`

	// **参数解释：** 合并请求所有E2E单号校验必须通过门禁是否通过。 **约束限制：** - true，合并请求所有E2E单号校验必须通过门禁通过。 - false，合并请求所有E2E单号校验必须通过门禁不通过。
	AllIssuesPassed *bool `json:"all_issues_passed,omitempty"`

	// **参数解释：** 合并请求只能关联一个单号门禁是否通过。 **约束限制：** - true，合并请求只能关联一个单号门禁通过。 - false，合并请求只能关联一个单号门禁不通过。
	OnlyOneIssuePassed *bool `json:"only_one_issue_passed,omitempty"`

	// **参数解释：** 合并请求检视门禁是否通过。 **约束限制：** - true，合并请求检视门禁通过。 - false，合并请求检视门禁不通过。
	ApprovalReviewersRequiredPassed *bool `json:"approval_reviewers_required_passed,omitempty"`

	// **参数解释：** 合并请求审核门禁是否通过。 **约束限制：** - true，合并请求审核门禁通过。 - false，合并请求审核门禁不通过。
	ApprovalApproversRequiredPassed *bool `json:"approval_approvers_required_passed,omitempty"`

	// **参数解释：** 合并请求星级评价门禁是否通过。 **约束限制：** - true，合并请求星级评价门禁通过。 - false，合并请求星级评价门禁不通过。
	EvaluationPassed *bool `json:"evaluation_passed,omitempty"`
	HttpStatusCode   int   `json:"-"`
}

func (o ShowMergeableStateOuterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeableStateOuterResponse struct{}"
	}

	return strings.Join([]string{"ShowMergeableStateOuterResponse", string(data)}, " ")
}
