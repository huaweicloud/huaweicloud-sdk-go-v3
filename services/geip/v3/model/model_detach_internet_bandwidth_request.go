package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachInternetBandwidthRequest Request Object
type DetachInternetBandwidthRequest struct {

	// global_eip_id
	GlobalEipId string `json:"global_eip_id"`

	// 是否强制解绑
	ForceUnbind *bool `json:"force_unbind,omitempty"`
}

func (o DetachInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DetachInternetBandwidthRequest", string(data)}, " ")
}
