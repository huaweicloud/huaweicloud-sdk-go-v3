package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportContentRequest Request Object
type ShowSecurityReportContentRequest struct {

	// **参数解释：** 报告ID，请从接口 “查询安全报告发送记录”（ListSecurityReportSendingRecords）中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ReportId string `json:"report_id"`

	// **参数解释：** 需要删除的订阅id，从“查询安全报告订阅列表”（ListSecurityReportSubscriptions）中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	SubscriptionId string `json:"subscription_id"`
}

func (o ShowSecurityReportContentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportContentRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportContentRequest", string(data)}, " ")
}
