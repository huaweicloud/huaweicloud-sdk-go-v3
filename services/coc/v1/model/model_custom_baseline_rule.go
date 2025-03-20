package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomBaselineRule 自定义基线规则
type CustomBaselineRule struct {

	// id
	Id *string `json:"id,omitempty"`

	// 规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 产品
	Product string `json:"product"`

	// 合规性报告等级
	ComplianceLevel string `json:"compliance_level"`

	// 基线补丁信息
	PatchItems *[]CustomBaselineRulePatchItem `json:"patch_items,omitempty"`
}

func (o CustomBaselineRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomBaselineRule struct{}"
	}

	return strings.Join([]string{"CustomBaselineRule", string(data)}, " ")
}
