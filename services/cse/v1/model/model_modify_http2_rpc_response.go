package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyHttp2RpcResponse Response Object
type ModifyHttp2RpcResponse struct {

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

func (o ModifyHttp2RpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHttp2RpcResponse struct{}"
	}

	return strings.Join([]string{"ModifyHttp2RpcResponse", string(data)}, " ")
}
