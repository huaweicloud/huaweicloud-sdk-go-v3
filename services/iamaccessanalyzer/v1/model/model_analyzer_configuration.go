package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalyzerConfiguration 分析器的配置项。
type AnalyzerConfiguration struct {
	UnusedAccess *AnalyzerConfigurationUnusedAccess `json:"unused_access,omitempty"`
}

func (o AnalyzerConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalyzerConfiguration struct{}"
	}

	return strings.Join([]string{"AnalyzerConfiguration", string(data)}, " ")
}
