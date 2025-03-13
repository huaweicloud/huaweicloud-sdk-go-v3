package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdentificationRuleResult 敏感识别规则诊断结果
type IdentificationRuleResult struct {
	Result *DiagnoseResult `json:"result,omitempty"`

	// 有风险的问题数量
	Count *int32 `json:"count,omitempty"`

	// 没有配置识别规则的密级列表
	AbnormalInfo *[]SecurityLevelInfo `json:"abnormal_info,omitempty"`
}

func (o IdentificationRuleResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentificationRuleResult struct{}"
	}

	return strings.Join([]string{"IdentificationRuleResult", string(data)}, " ")
}
