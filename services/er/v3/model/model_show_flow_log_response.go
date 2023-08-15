package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowLogResponse Response Object
type ShowFlowLogResponse struct {
	FlowLog *FlowLog `json:"flow_log,omitempty"`

	// 请求id
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowLogResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowLogResponse", string(data)}, " ")
}
