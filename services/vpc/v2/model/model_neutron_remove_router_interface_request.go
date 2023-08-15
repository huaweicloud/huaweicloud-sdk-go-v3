package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronRemoveRouterInterfaceRequest Request Object
type NeutronRemoveRouterInterfaceRequest struct {

	// 路由器ID
	RouterId string `json:"router_id"`

	Body *RouterInterfaceRequestBody `json:"body,omitempty"`
}

func (o NeutronRemoveRouterInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronRemoveRouterInterfaceRequest struct{}"
	}

	return strings.Join([]string{"NeutronRemoveRouterInterfaceRequest", string(data)}, " ")
}
