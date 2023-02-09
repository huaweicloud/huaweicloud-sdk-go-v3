package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowFlowLogRequest struct {

	// 流日志资源唯一标识
	FlowlogId string `json:"flowlog_id"`
}

func (o ShowFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowLogRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowLogRequest", string(data)}, " ")
}
