package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteRouterRequest Request Object
type NeutronDeleteRouterRequest struct {

	// 路由器ID
	RouterId string `json:"router_id"`
}

func (o NeutronDeleteRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteRouterRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteRouterRequest", string(data)}, " ")
}
