package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronDeleteNetworkRequest Request Object
type NeutronDeleteNetworkRequest struct {

	// 网络ID
	NetworkId string `json:"network_id"`
}

func (o NeutronDeleteNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronDeleteNetworkRequest struct{}"
	}

	return strings.Join([]string{"NeutronDeleteNetworkRequest", string(data)}, " ")
}
