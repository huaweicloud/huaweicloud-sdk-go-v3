package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCentralNetworkRequest Request Object
type ShowCentralNetworkRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`
}

func (o ShowCentralNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCentralNetworkRequest struct{}"
	}

	return strings.Join([]string{"ShowCentralNetworkRequest", string(data)}, " ")
}
