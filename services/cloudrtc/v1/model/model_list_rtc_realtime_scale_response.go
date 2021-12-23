package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcRealtimeScaleResponse struct {
	// 时间戳及相应时间的指标数值列表

	Scale *[]TimeValueData `json:"scale,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcRealtimeScaleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRealtimeScaleResponse struct{}"
	}

	return strings.Join([]string{"ListRtcRealtimeScaleResponse", string(data)}, " ")
}
