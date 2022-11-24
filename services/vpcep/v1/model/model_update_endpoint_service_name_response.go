package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEndpointServiceNameResponse struct {

	// 终端节点服务名称
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateEndpointServiceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceNameResponse", string(data)}, " ")
}
