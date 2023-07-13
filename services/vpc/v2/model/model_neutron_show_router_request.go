package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowRouterRequest Request Object
type NeutronShowRouterRequest struct {

	// 路由器ID
	RouterId string `json:"router_id"`
}

func (o NeutronShowRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowRouterRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowRouterRequest", string(data)}, " ")
}
