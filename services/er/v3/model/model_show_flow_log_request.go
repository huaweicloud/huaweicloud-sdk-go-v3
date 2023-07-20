package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlowLogRequest Request Object
type ShowFlowLogRequest struct {

	// 企业路由器实例ID
	ErId string `json:"er_id"`

	// 流日志ID
	FlowLogId string `json:"flow_log_id"`
}

func (o ShowFlowLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlowLogRequest struct{}"
	}

	return strings.Join([]string{"ShowFlowLogRequest", string(data)}, " ")
}
