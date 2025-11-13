package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityReportHistoryPeriodsRequest Request Object
type ListSecurityReportHistoryPeriodsRequest struct {

	// **参数解释：** 需要查询的订阅id，从“查询安全报告订阅列表”中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId string `json:"subscription_id"`

	// **参数解释：** 分页查询的单页返回数量，控制每次请求返回的记录条数。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页查询的起始位置，表示从第几条记录开始返回（从0开始计数）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSecurityReportHistoryPeriodsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityReportHistoryPeriodsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityReportHistoryPeriodsRequest", string(data)}, " ")
}
