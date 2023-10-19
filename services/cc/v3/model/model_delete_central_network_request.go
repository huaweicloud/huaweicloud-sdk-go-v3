package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCentralNetworkRequest Request Object
type DeleteCentralNetworkRequest struct {

	// 中心网络的ID。
	CentralNetworkId string `json:"central_network_id"`
}

func (o DeleteCentralNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCentralNetworkRequest struct{}"
	}

	return strings.Join([]string{"DeleteCentralNetworkRequest", string(data)}, " ")
}
