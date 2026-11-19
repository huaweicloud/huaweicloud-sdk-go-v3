package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityReportResponse Response Object
type ShowSecurityReportResponse struct {

	// **参数解释**： 报告ID **取值范围**： 最小值0，最大值2147483647
	ReportId *int32 `json:"report_id,omitempty"`

	// **参数解释**： 报告子ID **取值范围**： 最小值0，最大值2147483647
	ReportSubId *int32 `json:"report_sub_id,omitempty"`

	// **参数解释**： 报告名称 **取值范围**： 字符长度0-128位
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释**： 报告类别 **取值范围**： 字符长度0-64位
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释**： 报告发送的时间段 **取值范围**： 字符长度0-64位
	SendingPeriod *string `json:"sending_period,omitempty"`

	// **参数解释**： 主题URN **取值范围**： 字符长度0-256位
	TopicUrn *string `json:"topic_urn,omitempty"`

	// **参数解释**： 主题的显示名 **取值范围**： 字符长度0-256位
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 报告logo类型 **取值范围**： - default：华为logo - customize：自定义logo
	LogoType *string `json:"logo_type,omitempty"`

	// **参数解释**： 自定义logo的文件ID **取值范围**： 字符长度0-256位
	LogoId *string `json:"logo_id,omitempty"`

	// **参数解释**： 订阅的方式 **取值范围**： 字符长度0-64位
	SubscriptionType *string `json:"subscription_type,omitempty"`

	StatPeriod *StatPeriodRes `json:"stat_period,omitempty"`

	OverviewInfo *OverviewInfoRes `json:"overview_info,omitempty"`

	ScansInfo *ScansInfoRes `json:"scans_info,omitempty"`

	IntrusionInfo  *IntrusionInfoRes `json:"intrusion_info,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSecurityReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityReportResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityReportResponse", string(data)}, " ")
}
