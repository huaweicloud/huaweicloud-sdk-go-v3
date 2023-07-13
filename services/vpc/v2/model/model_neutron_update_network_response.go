package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateNetworkResponse Response Object
type NeutronUpdateNetworkResponse struct {
	Network        *NeutronNetwork `json:"network,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o NeutronUpdateNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateNetworkResponse struct{}"
	}

	return strings.Join([]string{"NeutronUpdateNetworkResponse", string(data)}, " ")
}
