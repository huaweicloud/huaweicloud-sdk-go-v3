package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateFlowLogRequest struct {

	// 流日志资源唯一标识
	FlowlogId string `json:"flowlog_id"`

	Body *UpdateFlowLogReqBody `json:"body,omitempty"`
}

func (o UpdateFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowLogRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlowLogRequest", string(data)}, " ")
}
