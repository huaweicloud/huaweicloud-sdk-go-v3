package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcRealtimeNetworkResponse Response Object
type ListRtcRealtimeNetworkResponse struct {

	// 时间戳及相应时间的指标数值列表
	Network *[]TimeDoubleValueData `json:"network,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcRealtimeNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcRealtimeNetworkResponse struct{}"
	}

	return strings.Join([]string{"ListRtcRealtimeNetworkResponse", string(data)}, " ")
}
