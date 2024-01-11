package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachVipBandwidthRequest Request Object
type AttachVipBandwidthRequest struct {

	// IPv6虚拟IP或者IPv6私网IP ID。
	VportId string `json:"vport_id"`

	Body *AttachVipBandwidthRequestBody `json:"body,omitempty"`
}

func (o AttachVipBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachVipBandwidthRequest struct{}"
	}

	return strings.Join([]string{"AttachVipBandwidthRequest", string(data)}, " ")
}
