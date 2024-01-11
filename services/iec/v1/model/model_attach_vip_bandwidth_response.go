package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachVipBandwidthResponse Response Object
type AttachVipBandwidthResponse struct {
	Vport          *VirtualPortResponse `json:"vport,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o AttachVipBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachVipBandwidthResponse struct{}"
	}

	return strings.Join([]string{"AttachVipBandwidthResponse", string(data)}, " ")
}
