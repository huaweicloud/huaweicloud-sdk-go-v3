package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlowLogResponse Response Object
type UpdateFlowLogResponse struct {
	FlowLog        *FlowLogResp `json:"flow_log,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowLogResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlowLogResponse", string(data)}, " ")
}
