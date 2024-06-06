package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalyzerUrn 分析器的唯一资源标识符。
type AnalyzerUrn struct {
}

func (o AnalyzerUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalyzerUrn struct{}"
	}

	return strings.Join([]string{"AnalyzerUrn", string(data)}, " ")
}
