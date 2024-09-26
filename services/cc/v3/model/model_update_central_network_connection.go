package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCentralNetworkConnection 更新中心网络连接信息。 * 带宽类型：按照带宽包分配、按照流量计费（当前版本暂不支持，预留后续支持按流量计费模式）、测试带宽类型 * 带宽包ID：按照带宽包分配必填，其他类型不填； * 带宽取值：按照带宽包分配和按照流量计费必填，代表限速带宽，按照带宽包分配下不能超过带宽包大小，测试带宽类型不填；
type UpdateCentralNetworkConnection struct {
	BandwidthType *BandwidthTypeEnum `json:"bandwidth_type"`

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	// 带宽值定义，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`
}

func (o UpdateCentralNetworkConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCentralNetworkConnection struct{}"
	}

	return strings.Join([]string{"UpdateCentralNetworkConnection", string(data)}, " ")
}
