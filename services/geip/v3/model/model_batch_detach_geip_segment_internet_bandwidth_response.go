package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDetachGeipSegmentInternetBandwidthResponse Response Object
type BatchDetachGeipSegmentInternetBandwidthResponse struct {

	// 异步任务，返回的任务ID
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDetachGeipSegmentInternetBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDetachGeipSegmentInternetBandwidthResponse struct{}"
	}

	return strings.Join([]string{"BatchDetachGeipSegmentInternetBandwidthResponse", string(data)}, " ")
}
