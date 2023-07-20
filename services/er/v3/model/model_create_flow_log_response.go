package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowLogResponse Response Object
type CreateFlowLogResponse struct {
	FlowLog *FlowLog `json:"flow_log,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	XClientToken   *string `json:"X-Client-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowLogResponse struct{}"
	}

	return strings.Join([]string{"CreateFlowLogResponse", string(data)}, " ")
}
