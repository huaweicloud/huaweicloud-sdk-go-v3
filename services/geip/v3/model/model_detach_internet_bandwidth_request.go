package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachInternetBandwidthRequest Request Object
type DetachInternetBandwidthRequest struct {
	GlobalEipId string `json:"global_eip_id"`

	ForceUnbind *bool `json:"force_unbind,omitempty"`
}

func (o DetachInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DetachInternetBandwidthRequest", string(data)}, " ")
}
