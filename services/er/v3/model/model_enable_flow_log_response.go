package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableFlowLogResponse Response Object
type EnableFlowLogResponse struct {
	FlowLog *FlowLog `json:"flow_log,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableFlowLogResponse struct{}"
	}

	return strings.Join([]string{"EnableFlowLogResponse", string(data)}, " ")
}
