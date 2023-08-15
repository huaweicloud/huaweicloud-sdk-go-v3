package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlowLogRequest Request Object
type UpdateFlowLogRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 流日志ID
	FlowLogId string `json:"flow_log_id"`

	Body *UpdateFlowLogRequestBody `json:"body,omitempty"`
}

func (o UpdateFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowLogRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlowLogRequest", string(data)}, " ")
}
