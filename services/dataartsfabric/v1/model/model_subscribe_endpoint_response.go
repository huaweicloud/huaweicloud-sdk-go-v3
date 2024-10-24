package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeEndpointResponse Response Object
type SubscribeEndpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SubscribeEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeEndpointResponse struct{}"
	}

	return strings.Join([]string{"SubscribeEndpointResponse", string(data)}, " ")
}
