package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Http2Rpc struct {

	// 网关的ID。
	GatewayId *string `json:"gatewayId,omitempty"`

	// 传递给插件的配置。
	PluginConfig *interface{} `json:"pluginConfig,omitempty"`

	Name *string `json:"name,omitempty"`

	// 路由的名称。
	RouteName *string `json:"routeName,omitempty"`

	// 目标路由的名称。
	RouteDestinationName *string `json:"routeDestinationName,omitempty"`

	Dubbo *Dubbo `json:"dubbo,omitempty"`
}

func (o Http2Rpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Http2Rpc struct{}"
	}

	return strings.Join([]string{"Http2Rpc", string(data)}, " ")
}
