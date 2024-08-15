package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailabilityZonesRequest Request Object
type ShowAvailabilityZonesRequest struct {

	// 区域ID。
	RegionId string `json:"region_id"`
}

func (o ShowAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"ShowAvailabilityZonesRequest", string(data)}, " ")
}
