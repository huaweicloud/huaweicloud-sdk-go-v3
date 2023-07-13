package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcRealtimeScaleDimensionResponse Response Object
type ListRtcRealtimeScaleDimensionResponse struct {

	// 维度分布信息
	Dimensions *[]RealtimeScaleDimensionValue `json:"dimensions,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcRealtimeScaleDimensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRealtimeScaleDimensionResponse struct{}"
	}

	return strings.Join([]string{"ListRtcRealtimeScaleDimensionResponse", string(data)}, " ")
}
