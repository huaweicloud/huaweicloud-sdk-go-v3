package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAnalyzerResponse Response Object
type ShowAnalyzerResponse struct {
	Analyzer       *AnalyzerSummary `json:"analyzer,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowAnalyzerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAnalyzerResponse struct{}"
	}

	return strings.Join([]string{"ShowAnalyzerResponse", string(data)}, " ")
}
