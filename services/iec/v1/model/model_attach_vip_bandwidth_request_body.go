package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachVipBandwidthRequestBody IPv6 IP绑定带宽请求体。
type AttachVipBandwidthRequestBody struct {

	// 带宽ID。
	BandwidthId string `json:"bandwidth_id"`
}

func (o AttachVipBandwidthRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachVipBandwidthRequestBody struct{}"
	}

	return strings.Join([]string{"AttachVipBandwidthRequestBody", string(data)}, " ")
}
