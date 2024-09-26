package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NonRequiredCentralNetworkPlaneId 中心网络平面ID。
type NonRequiredCentralNetworkPlaneId struct {

	// 中心网络平面ID。
	CentralNetworkPlaneId *string `json:"central_network_plane_id,omitempty"`
}

func (o NonRequiredCentralNetworkPlaneId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NonRequiredCentralNetworkPlaneId struct{}"
	}

	return strings.Join([]string{"NonRequiredCentralNetworkPlaneId", string(data)}, " ")
}
