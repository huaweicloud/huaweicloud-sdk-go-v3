package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInterRegionBandwidthRequest Request Object
type DeleteInterRegionBandwidthRequest struct {

	// 域间带宽实例ID。
	Id string `json:"id"`
}

func (o DeleteInterRegionBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInterRegionBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DeleteInterRegionBandwidthRequest", string(data)}, " ")
}
