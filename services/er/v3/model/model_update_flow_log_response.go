package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlowLogResponse Response Object
type UpdateFlowLogResponse struct {
	FlowLog *FlowLog `json:"flow_log,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowLogResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlowLogResponse", string(data)}, " ")
}
