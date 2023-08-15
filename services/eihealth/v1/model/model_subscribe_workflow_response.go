package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeWorkflowResponse Response Object
type SubscribeWorkflowResponse struct {

	// 流程id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SubscribeWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeWorkflowResponse struct{}"
	}

	return strings.Join([]string{"SubscribeWorkflowResponse", string(data)}, " ")
}
