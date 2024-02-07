package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachInternetBandwidthRequest Request Object
type BatchDetachInternetBandwidthRequest struct {
	Body *BatchDetachInternetBandwidthsGlobalEipRequestBody `json:"body,omitempty"`
}

func (o BatchDetachInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"BatchDetachInternetBandwidthRequest", string(data)}, " ")
}
