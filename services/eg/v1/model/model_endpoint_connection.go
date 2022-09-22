package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EndpointConnection struct {

	// 访问端点的服务ID
	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	// 访问端点的服务名称
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	// 访问端点的终端节点的报文ID
	MarkerId *string `json:"marker_id,omitempty"`

	// 访问端点的终端节点ID
	Id *string `json:"id,omitempty"`

	// 访问端点的终端节点IP
	Ip *string `json:"ip,omitempty"`

	// 访问端点的终端节点创建时间
	CreatedTime *string `json:"created_time,omitempty"`
}

func (o EndpointConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointConnection struct{}"
	}

	return strings.Join([]string{"EndpointConnection", string(data)}, " ")
}
