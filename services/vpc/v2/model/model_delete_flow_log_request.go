package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlowLogRequest Request Object
type DeleteFlowLogRequest struct {

	// 流日志资源唯一标识
	FlowlogId string `json:"flowlog_id"`
}

func (o DeleteFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlowLogRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlowLogRequest", string(data)}, " ")
}
