package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowLogResponse Response Object
type ShowFlowLogResponse struct {
	FlowLog        *FlowLogResp `json:"flow_log,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowLogResponse struct{}"
	}

	return strings.Join([]string{"ShowFlowLogResponse", string(data)}, " ")
}
