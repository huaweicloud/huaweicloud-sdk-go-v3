package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJunitCoverageSummaryResponse Response Object
type ListJunitCoverageSummaryResponse struct {
	Result *ListJunitCoverageSummaryResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJunitCoverageSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJunitCoverageSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListJunitCoverageSummaryResponse", string(data)}, " ")
}
