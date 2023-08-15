package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronShowNetworkRequest Request Object
type NeutronShowNetworkRequest struct {

	// 网络ID
	NetworkId string `json:"network_id"`
}

func (o NeutronShowNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronShowNetworkRequest struct{}"
	}

	return strings.Join([]string{"NeutronShowNetworkRequest", string(data)}, " ")
}
