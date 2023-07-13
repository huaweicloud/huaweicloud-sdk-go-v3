package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateRouterResponse Response Object
type NeutronCreateRouterResponse struct {
	Router         *NeutronRouter `json:"router,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o NeutronCreateRouterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateRouterResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateRouterResponse", string(data)}, " ")
}
