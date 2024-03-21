package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInternetBandwidthRequest Request Object
type DeleteInternetBandwidthRequest struct {
	InternetBandwidthId string `json:"internet_bandwidth_id"`
}

func (o DeleteInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DeleteInternetBandwidthRequest", string(data)}, " ")
}
