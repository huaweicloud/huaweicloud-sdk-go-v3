package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateRouterRequest Request Object
type NeutronUpdateRouterRequest struct {

	// 路由器ID
	RouterId string `json:"router_id"`

	Body *NeutronUpdateRouterRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateRouterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateRouterRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateRouterRequest", string(data)}, " ")
}
