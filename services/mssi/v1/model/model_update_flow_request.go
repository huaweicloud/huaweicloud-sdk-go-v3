package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlowRequest Request Object
type UpdateFlowRequest struct {

	// flow_id
	FlowId string `json:"flow_id"`

	Body *FlowMeta `json:"body,omitempty"`
}

func (o UpdateFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowRequest struct{}"
	}

	return strings.Join([]string{"UpdateFlowRequest", string(data)}, " ")
}
