package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnusedAnalysisRuleCriteria 未使用分析规则的条件。
type UnusedAnalysisRuleCriteria struct {

	// 账号ID列表。
	AccountIds *[]string `json:"account_ids,omitempty"`

	// 资源标签列表。
	ResourceTags *[]Tag `json:"resource_tags,omitempty"`
}

func (o UnusedAnalysisRuleCriteria) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedAnalysisRuleCriteria struct{}"
	}

	return strings.Join([]string{"UnusedAnalysisRuleCriteria", string(data)}, " ")
}
