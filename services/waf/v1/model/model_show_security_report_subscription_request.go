package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportSubscriptionRequest Request Object
type ShowSecurityReportSubscriptionRequest struct {

	// **参数解释：** 需要查询的订阅id，从“查询安全报告订阅列表”中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId string `json:"subscription_id"`
}

func (o ShowSecurityReportSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportSubscriptionRequest", string(data)}, " ")
}
