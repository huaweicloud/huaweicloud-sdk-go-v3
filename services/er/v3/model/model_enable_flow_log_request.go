package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableFlowLogRequest Request Object
type EnableFlowLogRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 流日志ID
	FlowLogId string `json:"flow_log_id"`
}

func (o EnableFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableFlowLogRequest struct{}"
	}

	return strings.Join([]string{"EnableFlowLogRequest", string(data)}, " ")
}
