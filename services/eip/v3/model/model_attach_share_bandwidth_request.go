package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachShareBandwidthRequest Request Object
type AttachShareBandwidthRequest struct {

	// 弹性公网ID
	PublicipId string `json:"publicip_id"`

	Body *AttachSharebwReq `json:"body,omitempty"`
}

func (o AttachShareBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachShareBandwidthRequest struct{}"
	}

	return strings.Join([]string{"AttachShareBandwidthRequest", string(data)}, " ")
}
