package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronUpdateNetworkRequest Request Object
type NeutronUpdateNetworkRequest struct {

	// 网络ID
	NetworkId string `json:"network_id"`

	Body *NeutronUpdateNetworkRequestBody `json:"body,omitempty"`
}

func (o NeutronUpdateNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateNetworkRequest struct{}"
	}

	return strings.Join([]string{"NeutronUpdateNetworkRequest", string(data)}, " ")
}
