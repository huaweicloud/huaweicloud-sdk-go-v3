package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkId 中心网络ID。
type CentralNetworkId struct {

	// 中心网络ID。
	CentralNetworkId string `json:"central_network_id"`
}

func (o CentralNetworkId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkId struct{}"
	}

	return strings.Join([]string{"CentralNetworkId", string(data)}, " ")
}
