package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlowLogsResponse Response Object
type ListFlowLogsResponse struct {

	// flow_log对象列表
	FlowLogs       *[]FlowLogResp `json:"flow_logs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListFlowLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlowLogsResponse struct{}"
	}

	return strings.Join([]string{"ListFlowLogsResponse", string(data)}, " ")
}
