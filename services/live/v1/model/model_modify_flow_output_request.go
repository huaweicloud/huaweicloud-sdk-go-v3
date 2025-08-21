package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyFlowOutputRequest Request Object
type ModifyFlowOutputRequest struct {

	// 流id
	FlowId string `json:"flow_id"`

	// 输出名称
	OutputName string `json:"output_name"`

	Body *UpdateFlowOutputRequestBody `json:"body,omitempty"`
}

func (o ModifyFlowOutputRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyFlowOutputRequest struct{}"
	}

	return strings.Join([]string{"ModifyFlowOutputRequest", string(data)}, " ")
}
