package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchLogsResponse Response Object
type ListSearchLogsResponse struct {
	AnalysisResults *AnalysisResults `json:"analysis_results,omitempty"`

	// 查询结果的条数
	Count *int64 `json:"count,omitempty"`

	// 返回的查询结果
	Results        *[]SearchResult `json:"results,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListSearchLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchLogsResponse struct{}"
	}

	return strings.Join([]string{"ListSearchLogsResponse", string(data)}, " ")
}
