package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowLogsResponse Response Object
type ListFlowLogsResponse struct {
	FlowLogs *[]FlowLog `json:"flow_logs,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	// 总计数量
	TotalCount *int64 `json:"total_count,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListFlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListFlowLogsResponse", string(data)}, " ")
}
