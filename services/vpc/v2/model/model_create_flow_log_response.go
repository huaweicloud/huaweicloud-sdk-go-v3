package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateFlowLogResponse struct {
	FlowLog        *FlowLogResp `json:"flow_log,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateFlowLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowLogResponse struct{}"
	}

	return strings.Join([]string{"CreateFlowLogResponse", string(data)}, " ")
}
