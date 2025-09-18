package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityReportHistoryPeriodsRequest Request Object
type ListSecurityReportHistoryPeriodsRequest struct {

	// 订阅ID
	SubscriptionId string `json:"subscription_id"`

	// 限制条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSecurityReportHistoryPeriodsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportHistoryPeriodsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityReportHistoryPeriodsRequest", string(data)}, " ")
}
