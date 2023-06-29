package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEndpointRequest Request Object
type CreateEndpointRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *Endpoint `json:"body,omitempty"`
}

func (o CreateEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateEndpointRequest", string(data)}, " ")
}
