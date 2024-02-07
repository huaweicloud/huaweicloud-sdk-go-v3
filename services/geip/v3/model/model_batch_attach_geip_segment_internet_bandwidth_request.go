package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachGeipSegmentInternetBandwidthRequest Request Object
type BatchAttachGeipSegmentInternetBandwidthRequest struct {
	Body *BatchAttachInternetBandwidthsGlobalEipSegmentRequestBody `json:"body,omitempty"`
}

func (o BatchAttachGeipSegmentInternetBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachGeipSegmentInternetBandwidthRequest struct{}"
	}

	return strings.Join([]string{"BatchAttachGeipSegmentInternetBandwidthRequest", string(data)}, " ")
}
