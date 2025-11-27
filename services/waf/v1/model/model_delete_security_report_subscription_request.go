package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityReportSubscriptionRequest Request Object
type DeleteSecurityReportSubscriptionRequest struct {

	// **参数解释：** 需要删除的订阅id，从“查询安全报告订阅列表”中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId string `json:"subscription_id"`
}

func (o DeleteSecurityReportSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityReportSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityReportSubscriptionRequest", string(data)}, " ")
}
