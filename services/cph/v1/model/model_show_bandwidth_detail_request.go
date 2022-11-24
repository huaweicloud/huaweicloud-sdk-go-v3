package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBandwidthDetailRequest struct {
}

func (o ShowBandwidthDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowBandwidthDetailRequest", string(data)}, " ")
}
