package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableFlowLogResponse Response Object
type DisableFlowLogResponse struct {
	FlowLog *FlowLog `json:"flow_log,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableFlowLogResponse struct{}"
	}

	return strings.Join([]string{"DisableFlowLogResponse", string(data)}, " ")
}
