package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteConnectionBandwidthSize 更改分支连接带宽大小的请求体。
type UpdateSiteConnectionBandwidthSize struct {

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`
}

func (o UpdateSiteConnectionBandwidthSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteConnectionBandwidthSize struct{}"
	}

	return strings.Join([]string{"UpdateSiteConnectionBandwidthSize", string(data)}, " ")
}
