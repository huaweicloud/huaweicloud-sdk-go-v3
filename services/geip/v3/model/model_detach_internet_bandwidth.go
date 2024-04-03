package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DetachInternetBandwidth struct {

	// 全域弹性公网IP段ID
	GlobalEipSegmentId string `json:"global_eip_segment_id"`
}

func (o DetachInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInternetBandwidth struct{}"
	}

	return strings.Join([]string{"DetachInternetBandwidth", string(data)}, " ")
}
