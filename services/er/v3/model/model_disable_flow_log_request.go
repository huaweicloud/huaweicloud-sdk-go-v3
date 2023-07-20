package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableFlowLogRequest Request Object
type DisableFlowLogRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 流日志ID
	FlowLogId string `json:"flow_log_id"`
}

func (o DisableFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableFlowLogRequest struct{}"
	}

	return strings.Join([]string{"DisableFlowLogRequest", string(data)}, " ")
}
