package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubnetBandwidthChangeOrderRequest Request Object
type CreateSubnetBandwidthChangeOrderRequest struct {

	// 云办公带宽id。
	BandwidthId string `json:"bandwidth_id"`

	Body *CreateSubnetBandwidthChangeOrderRequestBody `json:"body,omitempty"`
}

func (o CreateSubnetBandwidthChangeOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetBandwidthChangeOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateSubnetBandwidthChangeOrderRequest", string(data)}, " ")
}
