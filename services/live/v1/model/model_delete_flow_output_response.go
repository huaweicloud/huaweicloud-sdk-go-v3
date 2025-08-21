package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlowOutputResponse Response Object
type DeleteFlowOutputResponse struct {

	// 流ID
	FlowId *string `json:"flow_id,omitempty"`

	// 输出名称
	OutputName     *string `json:"output_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFlowOutputResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlowOutputResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlowOutputResponse", string(data)}, " ")
}
