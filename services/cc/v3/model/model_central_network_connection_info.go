package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CentralNetworkConnectionInfo struct {

	// 实例ID。
	Id string `json:"id"`

	// 中心网络平面ID。
	PlaneId string `json:"plane_id"`

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	// 带宽值定义，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`

	ConnectionType *ConnectionTypeEnum `json:"connection_type"`

	// 中心网络连接的两个端点定义，长度固定为2的数组。
	ConnectionPointPair []ConnectionPoint `json:"connection_point_pair"`

	State *CentralNetworkConnectionStateEnum `json:"state"`
}

func (o CentralNetworkConnectionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkConnectionInfo struct{}"
	}

	return strings.Join([]string{"CentralNetworkConnectionInfo", string(data)}, " ")
}
