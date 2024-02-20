package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnalyzerRequest Request Object
type ShowAnalyzerRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`
}

func (o ShowAnalyzerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalyzerRequest struct{}"
	}

	return strings.Join([]string{"ShowAnalyzerRequest", string(data)}, " ")
}
