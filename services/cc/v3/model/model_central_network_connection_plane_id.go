package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CentralNetworkConnectionPlaneId 中心网络平面ID。
type CentralNetworkConnectionPlaneId struct {

	// 中心网络平面ID。
	PlaneId string `json:"plane_id"`
}

func (o CentralNetworkConnectionPlaneId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkConnectionPlaneId struct{}"
	}

	return strings.Join([]string{"CentralNetworkConnectionPlaneId", string(data)}, " ")
}
