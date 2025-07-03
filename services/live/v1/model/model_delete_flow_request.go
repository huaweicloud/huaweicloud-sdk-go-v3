package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlowRequest Request Object
type DeleteFlowRequest struct {

	// 流id
	FlowId string `json:"flow_id"`
}

func (o DeleteFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlowRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlowRequest", string(data)}, " ")
}
