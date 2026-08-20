package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessInstanceResponseResultReviewConfig 评审配置
type ProcessInstanceResponseResultReviewConfig struct {

	// 审批类型
	ApprovalType *int32 `json:"approval_type,omitempty"`

	// 审批进度
	RatioValue *string `json:"ratio_value,omitempty"`

	// 是否跳过决策
	SkipDecisioning *bool `json:"skip_decisioning,omitempty"`

	// 决策角色
	ApprovalRoles *string `json:"approval_roles,omitempty"`

	// 审批角色
	ReviewRoles *string `json:"review_roles,omitempty"`
}

func (o ProcessInstanceResponseResultReviewConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResultReviewConfig struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResultReviewConfig", string(data)}, " ")
}
