package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkDefaultPlaneId 中心网络默认平面的ID。
type CentralNetworkDefaultPlaneId struct {

	// 中心网络默认平面的ID。
	DefaultPlaneId string `json:"default_plane_id"`
}

func (o CentralNetworkDefaultPlaneId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkDefaultPlaneId struct{}"
	}

	return strings.Join([]string{"CentralNetworkDefaultPlaneId", string(data)}, " ")
}
