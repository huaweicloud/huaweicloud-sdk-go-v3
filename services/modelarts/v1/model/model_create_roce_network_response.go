package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoceNetworkResponse Response Object
type CreateRoceNetworkResponse struct {
	Network        *ServerRoceNetwork `json:"network,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o CreateRoceNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoceNetworkResponse struct{}"
	}

	return strings.Join([]string{"CreateRoceNetworkResponse", string(data)}, " ")
}
