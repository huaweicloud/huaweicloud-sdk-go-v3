package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcHistoryScaleResponse struct {
	// 时间戳及相应时间的指标数值

	Scale *[]RtcHistoryScaleTimeValue `json:"scale,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcHistoryScaleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryScaleResponse struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryScaleResponse", string(data)}, " ")
}
