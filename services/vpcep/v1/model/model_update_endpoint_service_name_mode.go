package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointServiceNameMode 修改终端节点服务名称
type UpdateEndpointServiceNameMode struct {

	// 终端节点服务名称
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`
}

func (o UpdateEndpointServiceNameMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointServiceNameMode struct{}"
	}

	return strings.Join([]string{"UpdateEndpointServiceNameMode", string(data)}, " ")
}
