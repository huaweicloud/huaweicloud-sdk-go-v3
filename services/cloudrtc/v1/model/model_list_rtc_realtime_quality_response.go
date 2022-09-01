package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcRealtimeQualityResponse struct {

	// 时间戳及相应时间的指标数值列表
	Quality *[]TimeDoubleValueData `json:"quality,omitempty" xml:"quality"`

	XRequestId     *string `json:"X-request-id,omitempty" xml:"X-request-id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcRealtimeQualityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRealtimeQualityResponse struct{}"
	}

	return strings.Join([]string{"ListRtcRealtimeQualityResponse", string(data)}, " ")
}
