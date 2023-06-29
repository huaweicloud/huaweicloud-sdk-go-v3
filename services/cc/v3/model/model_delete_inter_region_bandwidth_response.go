package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInterRegionBandwidthResponse Response Object
type DeleteInterRegionBandwidthResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInterRegionBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInterRegionBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DeleteInterRegionBandwidthResponse", string(data)}, " ")
}
