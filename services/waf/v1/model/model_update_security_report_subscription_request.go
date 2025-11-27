package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityReportSubscriptionRequest Request Object
type UpdateSecurityReportSubscriptionRequest struct {

	// **参数解释：** 需要删除的订阅id，从“查询安全报告订阅列表”中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId string `json:"subscription_id"`

	Body *UpdateSecurityReportSubscriptionRequestBody `json:"body,omitempty"`
}

func (o UpdateSecurityReportSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityReportSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityReportSubscriptionRequest", string(data)}, " ")
}
