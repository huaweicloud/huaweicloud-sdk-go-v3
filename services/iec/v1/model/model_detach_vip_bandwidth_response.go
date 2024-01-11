package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachVipBandwidthResponse Response Object
type DetachVipBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DetachVipBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachVipBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DetachVipBandwidthResponse", string(data)}, " ")
}
