package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlowOutputRequest Request Object
type DeleteFlowOutputRequest struct {

	// 流id
	FlowId string `json:"flow_id"`

	// 输出名称
	OutputName string `json:"output_name"`
}

func (o DeleteFlowOutputRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlowOutputRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlowOutputRequest", string(data)}, " ")
}
