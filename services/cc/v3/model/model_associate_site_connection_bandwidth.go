package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateSiteConnectionBandwidth struct {

	// 全域互联带宽ID。
	GlobalConnectionBandwidthId *string `json:"global_connection_bandwidth_id,omitempty"`

	// 带宽值，单位Mbps。
	BandwidthSize *int64 `json:"bandwidth_size,omitempty"`
}

func (o AssociateSiteConnectionBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSiteConnectionBandwidth struct{}"
	}

	return strings.Join([]string{"AssociateSiteConnectionBandwidth", string(data)}, " ")
}
