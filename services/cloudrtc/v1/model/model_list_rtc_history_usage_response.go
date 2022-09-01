package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcHistoryUsageResponse struct {

	// 时间戳及相应时间的指标数值列表
	Usage *[]RtcHistoryUsage `json:"usage,omitempty" xml:"usage"`

	XRequestId     *string `json:"X-request-id,omitempty" xml:"X-request-id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcHistoryUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryUsageResponse struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryUsageResponse", string(data)}, " ")
}
