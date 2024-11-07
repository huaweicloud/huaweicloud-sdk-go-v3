package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

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
