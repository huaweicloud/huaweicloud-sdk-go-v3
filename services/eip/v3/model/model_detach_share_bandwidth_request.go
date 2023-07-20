package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetachShareBandwidthRequest Request Object
type DetachShareBandwidthRequest struct {

	// 弹性公网ID
	PublicipId string `json:"publicip_id"`

	Body *DetachSharedbwReq `json:"body,omitempty"`
}

func (o DetachShareBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachShareBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DetachShareBandwidthRequest", string(data)}, " ")
}
