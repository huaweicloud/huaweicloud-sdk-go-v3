package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointServiceNameRequestBody 修改终端节点服务名称
type UpdateEndpointServiceNameRequestBody struct {

	// 终端节点服务名称
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`
}

func (o UpdateEndpointServiceNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceNameRequestBody", string(data)}, " ")
}
