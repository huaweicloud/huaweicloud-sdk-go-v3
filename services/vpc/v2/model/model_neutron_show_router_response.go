package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowRouterResponse Response Object
type NeutronShowRouterResponse struct {
	Router         *NeutronRouter `json:"router,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o NeutronShowRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowRouterResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowRouterResponse", string(data)}, " ")
}
