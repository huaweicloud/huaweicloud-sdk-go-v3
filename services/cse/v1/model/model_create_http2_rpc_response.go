package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttp2RpcResponse Response Object
type CreateHttp2RpcResponse struct {

	// 网关的ID。
	GatewayId *string `json:"gatewayId,omitempty"`

	// 传递给插件的配置。
	PluginConfig *interface{} `json:"pluginConfig,omitempty"`

	Name *string `json:"name,omitempty"`

	// 路由的名称。
	RouteName *string `json:"routeName,omitempty"`

	// 目标路由的名称。
	RouteDestinationName *string `json:"routeDestinationName,omitempty"`

	Dubbo          *Dubbo `json:"dubbo,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateHttp2RpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttp2RpcResponse struct{}"
	}

	return strings.Join([]string{"CreateHttp2RpcResponse", string(data)}, " ")
}
