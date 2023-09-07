package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlowResponse Response Object
type UpdateFlowResponse struct {

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

func (o UpdateFlowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlowResponse struct{}"
	}

	return strings.Join([]string{"UpdateFlowResponse", string(data)}, " ")
}
