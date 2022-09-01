package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcHistoryScaleResponse struct {

	// 时间戳及相应时间的指标数值列表
	Scale *[]RtcHistoryScaleTimeValue `json:"scale,omitempty" xml:"scale"`

	XRequestId     *string `json:"X-request-id,omitempty" xml:"X-request-id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcHistoryScaleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryScaleResponse struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryScaleResponse", string(data)}, " ")
}
