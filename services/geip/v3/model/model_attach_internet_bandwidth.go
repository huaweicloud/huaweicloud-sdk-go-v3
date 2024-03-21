package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachInternetBandwidth struct {

	// 全域弹性公网IP段的ID
	GlobalEipSegmentId string `json:"global_eip_segment_id"`

	// 全域公网带宽的ID
	InternetBandwidthId string `json:"internet_bandwidth_id"`
}

func (o AttachInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternetBandwidth struct{}"
	}

	return strings.Join([]string{"AttachInternetBandwidth", string(data)}, " ")
}
