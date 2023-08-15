package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronAddRouterInterfaceRequest Request Object
type NeutronAddRouterInterfaceRequest struct {

	// 路由器ID
	RouterId string `json:"router_id"`

	Body *RouterInterfaceRequestBody `json:"body,omitempty"`
}

func (o NeutronAddRouterInterfaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronAddRouterInterfaceRequest struct{}"
	}

	return strings.Join([]string{"NeutronAddRouterInterfaceRequest", string(data)}, " ")
}
