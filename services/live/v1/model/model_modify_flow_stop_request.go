package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyFlowStopRequest Request Object
type ModifyFlowStopRequest struct {

	// 流id
	FlowId string `json:"flow_id"`
}

func (o ModifyFlowStopRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowStopRequest struct{}"
	}

	return strings.Join([]string{"ModifyFlowStopRequest", string(data)}, " ")
}
