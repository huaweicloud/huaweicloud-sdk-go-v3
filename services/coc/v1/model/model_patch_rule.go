package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchRule 操作系统补丁批准过滤规则
type PatchRule struct {

	// 批准规则过滤
	PatchFilters []PatchFilter `json:"patch_filters"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 合规性报告级别
	ComplianceLevel string `json:"compliance_level"`

	// 指定天数后批准补丁，指定的天数
	ApproveAfterDays *int32 `json:"approve_after_days,omitempty"`

	// 批准指定日期之前发布的补丁,指定的日期时间戳
	ApproveUntilDate *int64 `json:"approve_until_date,omitempty"`

	// 是否包括非安全性更新
	EnableNonSecurity bool `json:"enable_non_security"`
}

func (o PatchRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchRule struct{}"
	}

	return strings.Join([]string{"PatchRule", string(data)}, " ")
}
