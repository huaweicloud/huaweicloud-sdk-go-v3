package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowOutputRequest Request Object
type CreateFlowOutputRequest struct {

	// 流id
	FlowId string `json:"flow_id"`

	Body *[]AddFlowOutputsRequest `json:"body,omitempty"`
}

func (o CreateFlowOutputRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowOutputRequest struct{}"
	}

	return strings.Join([]string{"CreateFlowOutputRequest", string(data)}, " ")
}
