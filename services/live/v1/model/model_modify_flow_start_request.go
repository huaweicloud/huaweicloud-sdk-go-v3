package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyFlowStartRequest Request Object
type ModifyFlowStartRequest struct {

	// 流id
	FlowId string `json:"flow_id"`
}

func (o ModifyFlowStartRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowStartRequest struct{}"
	}

	return strings.Join([]string{"ModifyFlowStartRequest", string(data)}, " ")
}
