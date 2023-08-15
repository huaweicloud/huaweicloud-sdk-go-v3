package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInterRegionBandwidthRequest Request Object
type UpdateInterRegionBandwidthRequest struct {

	// 域间带宽实例ID。
	Id string `json:"id"`

	Body *UpdateInterRegionBandwidthRequestBody `json:"body,omitempty"`
}

func (o UpdateInterRegionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInterRegionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdateInterRegionBandwidthRequest", string(data)}, " ")
}
