package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkPlaneId 中心网络平面ID。
type CentralNetworkPlaneId struct {

	// 中心网络平面ID。
	CentralNetworkPlaneId string `json:"central_network_plane_id"`
}

func (o CentralNetworkPlaneId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkPlaneId struct{}"
	}

	return strings.Join([]string{"CentralNetworkPlaneId", string(data)}, " ")
}
