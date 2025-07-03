package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlowResponse Response Object
type DeleteFlowResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 流id
	FlowId         *string `json:"flow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlowResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlowResponse", string(data)}, " ")
}
