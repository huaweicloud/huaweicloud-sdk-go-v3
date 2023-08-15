package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowNetworkResponse Response Object
type NeutronShowNetworkResponse struct {
	Network        *NeutronNetwork `json:"network,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o NeutronShowNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowNetworkResponse struct{}"
	}

	return strings.Join([]string{"NeutronShowNetworkResponse", string(data)}, " ")
}
