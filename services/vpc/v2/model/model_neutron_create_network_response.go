package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreateNetworkResponse Response Object
type NeutronCreateNetworkResponse struct {
	Network        *NeutronNetwork `json:"network,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o NeutronCreateNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateNetworkResponse struct{}"
	}

	return strings.Join([]string{"NeutronCreateNetworkResponse", string(data)}, " ")
}
