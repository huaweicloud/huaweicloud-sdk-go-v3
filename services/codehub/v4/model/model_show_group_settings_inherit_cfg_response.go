package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupSettingsInheritCfgResponse Response Object
type ShowGroupSettingsInheritCfgResponse struct {

	// **参数解释：** 是否可用。
	CanUpdate *bool `json:"can_update,omitempty"`

	// **参数解释：** 记录id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 项目id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释：** 代码组id。
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// **参数解释：** 父id。
	ParentId *int32 `json:"parent_id,omitempty"`

	// **参数解释：** 排序id。
	Ownership *int32 `json:"ownership,omitempty"`

	// **参数解释：** 排序id。
	Pbi *int32 `json:"pbi,omitempty"`

	// **参数解释：** 保护分支策略。
	ProtectedBranches *int32 `json:"protected_branches,omitempty"`

	// **参数解释：** 保护tag策略。
	ProtectedTags *int32 `json:"protected_tags,omitempty"`

	// **参数解释：** 推送规则策略。
	PushRules *int32 `json:"push_rules,omitempty"`

	// **参数解释：** cr策略。
	ChangeRequests *int32 `json:"change_requests,omitempty"`

	// **参数解释：** 排序id。
	CustomCtrlItems *int32 `json:"custom_ctrl_items,omitempty"`

	// **参数解释：** 排序id。
	Reviews *int32 `json:"reviews,omitempty"`

	// **参数解释：** issue继承模式。
	Issues *int32 `json:"issues,omitempty"`

	// **参数解释：** 排序id。
	CrEvaluation *int32 `json:"cr_evaluation,omitempty"`

	// **参数解释：** e2e策略。
	E2eSettings *int32 `json:"e2e_settings,omitempty"`

	// **参数解释：** 提交策略。
	CommitterSettings *int32 `json:"committer_settings,omitempty"`

	// **参数解释：** webhook策略。
	WebhookSettings *int32 `json:"webhook_settings,omitempty"`

	// **参数解释：** 排序id。
	StreamEventSettings *int32 `json:"stream_event_settings,omitempty"`

	// **参数解释：** 排序id。
	PipelineSettings *int32 `json:"pipeline_settings,omitempty"`

	// **参数解释：** issue模板继承模式。
	IssueTemplates *int32 `json:"issue_templates,omitempty"`

	// **参数解释：** 排序id。
	CrCommentEmplates *int32 `json:"cr_comment_emplates,omitempty"`

	// **参数解释：** 排序id。
	MergeRequests *int32 `json:"merge_requests,omitempty"`

	// **参数解释：** 合并请求策略。
	MrBranchPolicies *int32 `json:"mr_branch_policies,omitempty"`

	// **参数解释：** 仓库策略。
	RepositorySettings *int32 `json:"repository_settings,omitempty"`

	// **参数解释：** 部署秘钥策略。
	DeployKeys *int32 `json:"deploy_keys,omitempty"`

	// **参数解释：** 水印策略。
	Watermark *int32 `json:"watermark,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdateAt       *string `json:"update_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGroupSettingsInheritCfgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupSettingsInheritCfgResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupSettingsInheritCfgResponse", string(data)}, " ")
}
