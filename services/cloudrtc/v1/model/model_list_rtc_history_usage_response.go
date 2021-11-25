package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcHistoryUsageResponse struct {
	// 时间戳及相应时间的指标数值

	Usage *[]RtcHistoryUsage `json:"usage,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcHistoryUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryUsageResponse struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryUsageResponse", string(data)}, " ")
}
