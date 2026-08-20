package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeEndpointProxyRequest Request Object
type InvokeEndpointProxyRequest struct {
	Body *EndpointProxyParam `json:"body,omitempty"`
}

func (o InvokeEndpointProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeEndpointProxyRequest struct{}"
	}

	return strings.Join([]string{"InvokeEndpointProxyRequest", string(data)}, " ")
}
