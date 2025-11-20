package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnusedAnalysisRule 未使用分析规则。
type UnusedAnalysisRule struct {

	// 排除规则。
	Exclusions *[]UnusedAnalysisRuleCriteria `json:"exclusions,omitempty"`
}

func (o UnusedAnalysisRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedAnalysisRule struct{}"
	}

	return strings.Join([]string{"UnusedAnalysisRule", string(data)}, " ")
}
