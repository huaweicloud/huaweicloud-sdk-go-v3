package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInterRegionBandwidthRequest Request Object
type ShowInterRegionBandwidthRequest struct {

	// 实例ID。
	Id string `json:"id"`
}

func (o ShowInterRegionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInterRegionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"ShowInterRegionBandwidthRequest", string(data)}, " ")
}
