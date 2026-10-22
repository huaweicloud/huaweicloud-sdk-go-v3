package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiGuardOverviewSummaryResponse Response Object
type ListAiGuardOverviewSummaryResponse struct {

	// 数据条数
	Total *int32 `json:"total,omitempty"`

	// 总览统计数据列表
	Items *[]OverviewStatistics `json:"items,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAiGuardOverviewSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiGuardOverviewSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListAiGuardOverviewSummaryResponse", string(data)}, " ")
}
