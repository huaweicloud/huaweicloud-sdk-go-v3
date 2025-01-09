package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubnetBandwidthRequest Request Object
type DeleteSubnetBandwidthRequest struct {

	// 云办公带宽id。
	BandwidthId string `json:"bandwidth_id"`
}

func (o DeleteSubnetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubnetBandwidthRequest", string(data)}, " ")
}
