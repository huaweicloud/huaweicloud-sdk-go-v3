package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAnalyzersResponse Response Object
type ListAnalyzersResponse struct {

	// 分析器列表信息。
	Analyzers *[]AnalyzerSummary `json:"analyzers,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAnalyzersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAnalyzersResponse struct{}"
	}

	return strings.Join([]string{"ListAnalyzersResponse", string(data)}, " ")
}
