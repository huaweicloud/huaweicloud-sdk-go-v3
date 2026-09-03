package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlertConfigVo struct {

	// 告警表达式
	AlertExpression *[]AlertExpression `json:"alert_expression,omitempty"`

	// 告警区间，开始时间
	AlertPeriodBegin *string `json:"alertPeriodBegin,omitempty"`

	// 告警区间，开始时间
	AlertPeriodEnd *string `json:"alertPeriodEnd,omitempty"`

	BlockAlert *BlockAlert `json:"blockAlert,omitempty"`

	DefaultAlertTemplate *AlertTemplate `json:"defaultAlertTemplate,omitempty"`

	// 告警开启 0关闭 1开启
	Enable *string `json:"enable,omitempty"`

	ErrorAlert *ErrorAlert `json:"errorAlert,omitempty"`

	FailedAlert *FailedAlert `json:"failedAlert,omitempty"`

	// 告警恢复通知开关 0关闭 1开启
	RecoverNoticeEnable *string `json:"recoverNoticeEnable,omitempty"`

	// 告警收敛开关 0关闭 1开启
	RestrainAlertEnable *string `json:"restrainAlertEnable,omitempty"`

	// 智能告警 成功多少次后发送恢复告警
	ResumeAlertNum *int32 `json:"resume_alert_num,omitempty"`

	// 智能告警 指定时间后发送恢复告警
	ResumeAlertTime *string `json:"resumeAlertTime,omitempty"`

	TimeoutAlert *TimeoutAlert `json:"timeoutAlert,omitempty"`

	TimeoutAlertV4 *TimeoutAlert `json:"timeoutAlertV4,omitempty"`
}

func (o AlertConfigVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertConfigVo struct{}"
	}

	return strings.Join([]string{"AlertConfigVo", string(data)}, " ")
}
