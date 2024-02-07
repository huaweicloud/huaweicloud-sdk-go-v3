package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateInternetBandwidthRequest Request Object
type BatchCreateInternetBandwidthRequest struct {
	Body *BatchCreateInternetBandwidthRequestBody `json:"body,omitempty"`
}

func (o BatchCreateInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateInternetBandwidthRequest", string(data)}, " ")
}
