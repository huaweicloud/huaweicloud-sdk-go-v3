package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubnetBandwidthControlListRequest Request Object
type UpdateSubnetBandwidthControlListRequest struct {

	// 云办公带宽id。
	BandwidthId string `json:"bandwidth_id"`

	Body *UpdateSubnetBandwidthControlListReq `json:"body,omitempty"`
}

func (o UpdateSubnetBandwidthControlListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetBandwidthControlListRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubnetBandwidthControlListRequest", string(data)}, " ")
}
