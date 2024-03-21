package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInternetBandwidthRequest Request Object
type AttachInternetBandwidthRequest struct {
	GlobalEipId string `json:"global_eip_id"`

	Body *AttachInternetBandwidthGlobalEipRequestBody `json:"body,omitempty"`
}

func (o AttachInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"AttachInternetBandwidthRequest", string(data)}, " ")
}
