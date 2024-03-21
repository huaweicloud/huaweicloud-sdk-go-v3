package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInternetBandwidthRequest Request Object
type UpdateInternetBandwidthRequest struct {
	InternetBandwidthId string `json:"internet_bandwidth_id"`

	Body *UpdateInternetBandwidthRequestBody `json:"body,omitempty"`
}

func (o UpdateInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"UpdateInternetBandwidthRequest", string(data)}, " ")
}
