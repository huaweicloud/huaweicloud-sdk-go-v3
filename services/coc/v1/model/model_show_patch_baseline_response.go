package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPatchBaselineResponse Response Object
type ShowPatchBaselineResponse struct {

	// 补丁基准ID
	BaselineId *string `json:"baseline_id,omitempty"`

	// 补丁基准名称
	Name *string `json:"name,omitempty"`

	// 补丁基准描述
	Description *string `json:"description,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 基线使用场景(CCE、ECS、BMS)
	BaselineScene *string `json:"baseline_scene,omitempty"`

	// 基线规则类型(Standard,Custom)
	RuleType *string `json:"rule_type,omitempty"`

	// 操作系统
	OperatingSystem *string `json:"operating_system,omitempty"`

	// 补丁基准类型（公共/自定义）
	Type *string `json:"type,omitempty"`

	// 操作系统的批准规则
	ApprovalRules *[]PatchRule `json:"approval_rules,omitempty"`

	// 自定义基线列表
	CustomBaselineRules *[]CustomBaselineRule `json:"custom_baseline_rules,omitempty"`

	// 是否默认基准
	DefaultPatchBaseline *bool `json:"default_patch_baseline,omitempty"`

	// 拒绝的补丁
	RejectedPatches *string `json:"rejected_patches,omitempty"`

	// 拒绝策略
	RejectedAction *string `json:"rejected_action,omitempty"`

	// 批准的补丁
	ApprovedPatches *string `json:"approved_patches,omitempty"`

	// 批准的补丁合规性级别
	ApprovedCompliance *string `json:"approved_compliance,omitempty"`

	// 批准的补丁是否安全更新
	EnableSecurityApproved *bool `json:"enable_security_approved,omitempty"`

	// 创建时间
	CreatedTime *int64 `json:"created_time,omitempty"`

	// modifiedTime
	ModifiedTime   *int64 `json:"modified_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowPatchBaselineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPatchBaselineResponse struct{}"
	}

	return strings.Join([]string{"ShowPatchBaselineResponse", string(data)}, " ")
}
