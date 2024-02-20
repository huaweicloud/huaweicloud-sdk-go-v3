package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAnalyzerRequest Request Object
type DeleteAnalyzerRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`
}

func (o DeleteAnalyzerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnalyzerRequest struct{}"
	}

	return strings.Join([]string{"DeleteAnalyzerRequest", string(data)}, " ")
}
