package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteNetworkResponse Response Object
type NeutronDeleteNetworkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NeutronDeleteNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteNetworkResponse struct{}"
	}

	return strings.Join([]string{"NeutronDeleteNetworkResponse", string(data)}, " ")
}
