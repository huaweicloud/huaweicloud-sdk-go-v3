package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachGeipSegmentInternetBandwidthRequest Request Object
type BatchDetachGeipSegmentInternetBandwidthRequest struct {
	Body *BatchDetachInternetBandwidthsGlobalEipSegmentRequestBody `json:"body,omitempty"`
}

func (o BatchDetachGeipSegmentInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachGeipSegmentInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"BatchDetachGeipSegmentInternetBandwidthRequest", string(data)}, " ")
}
