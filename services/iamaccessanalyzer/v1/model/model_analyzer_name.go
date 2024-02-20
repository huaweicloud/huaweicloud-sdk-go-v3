package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnalyzerName 分析器的名称。
type AnalyzerName struct {
}

func (o AnalyzerName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalyzerName struct{}"
	}

	return strings.Join([]string{"AnalyzerName", string(data)}, " ")
}
