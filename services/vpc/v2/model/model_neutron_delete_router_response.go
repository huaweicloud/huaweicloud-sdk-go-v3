package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteRouterResponse Response Object
type NeutronDeleteRouterResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteRouterResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteRouterResponse", string(data)}, " ")
}
