package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachVipBandwidthRequest Request Object
type DetachVipBandwidthRequest struct {

	// IPv6虚拟IP或者IPv6私网IP ID。
	VportId string `json:"vport_id"`
}

func (o DetachVipBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachVipBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DetachVipBandwidthRequest", string(data)}, " ")
}
