package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetBandwidthControlListResponse Response Object
type UpdateSubnetBandwidthControlListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSubnetBandwidthControlListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetBandwidthControlListResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubnetBandwidthControlListResponse", string(data)}, " ")
}
