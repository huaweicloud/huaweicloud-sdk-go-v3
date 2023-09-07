package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowResponse Response Object
type CreateFlowResponse struct {

	// api流注册到apig的url
	ApigUrl *string `json:"apig_url,omitempty"`

	// ID
	FlowId *string `json:"flow_id,omitempty"`

	// 函数连接器对应functionId
	Steps *[]interface{} `json:"steps,omitempty"`

	// webhook触发url的ID
	Webhook        *string `json:"webhook,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowResponse struct{}"
	}

	return strings.Join([]string{"CreateFlowResponse", string(data)}, " ")
}
