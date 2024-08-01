package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalyzerConfigurationUnusedAccess 未使用的访问分析器的配置项。
type AnalyzerConfigurationUnusedAccess struct {

	// 生成分析结果的预设天数。
	UnusedAccessAge *int32 `json:"unused_access_age,omitempty"`
}

func (o AnalyzerConfigurationUnusedAccess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalyzerConfigurationUnusedAccess struct{}"
	}

	return strings.Join([]string{"AnalyzerConfigurationUnusedAccess", string(data)}, " ")
}
