package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteConnectionBandwidth 更改分支连接带宽包的请求体。
type UpdateSiteConnectionBandwidth struct {

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`
}

func (o UpdateSiteConnectionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteConnectionBandwidth struct{}"
	}

	return strings.Join([]string{"UpdateSiteConnectionBandwidth", string(data)}, " ")
}
