package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowOutputResponse Response Object
type CreateFlowOutputResponse struct {

	// 流ID
	FlowId *string `json:"flow_id,omitempty"`

	// 所有输出信息
	Outputs        *[]FlowsOutput `json:"outputs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateFlowOutputResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowOutputResponse struct{}"
	}

	return strings.Join([]string{"CreateFlowOutputResponse", string(data)}, " ")
}
