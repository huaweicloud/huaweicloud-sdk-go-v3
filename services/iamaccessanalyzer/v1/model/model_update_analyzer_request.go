package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAnalyzerRequest Request Object
type UpdateAnalyzerRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *UpdateAnalyzerReqBody `json:"body,omitempty"`
}

func (o UpdateAnalyzerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnalyzerRequest struct{}"
	}

	return strings.Join([]string{"UpdateAnalyzerRequest", string(data)}, " ")
}
