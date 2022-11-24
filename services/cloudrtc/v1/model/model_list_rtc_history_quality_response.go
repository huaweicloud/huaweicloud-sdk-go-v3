package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcHistoryQualityResponse struct {

	// 时间戳及相应时间的指标数值列表
	Quality *[]RtcHistoryQualityTimeValue `json:"quality,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcHistoryQualityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcHistoryQualityResponse struct{}"
	}

	return strings.Join([]string{"ListRtcHistoryQualityResponse", string(data)}, " ")
}
