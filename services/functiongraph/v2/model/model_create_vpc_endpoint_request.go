package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpcEndpointRequest Request Object
type CreateVpcEndpointRequest struct {

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *CreateVpcEndpointRequestBody `json:"body,omitempty"`
}

func (o CreateVpcEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateVpcEndpointRequest", string(data)}, " ")
}
