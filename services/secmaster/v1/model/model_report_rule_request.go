package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReportRuleRequest 创建报告发送规则请求体
type ReportRuleRequest struct {

	// 报告发送任务cron表达式
	Rule *string `json:"rule,omitempty"`

	// 报告发送任务启动截止时间
	RuleEnd *string `json:"rule_end,omitempty"`

	StartTime *ReportRuleRequestStartTime `json:"start_time,omitempty"`

	EndTime *ReportRuleRequestEndTime `json:"end_time,omitempty"`
}

func (o ReportRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReportRuleRequest struct{}"
	}

	return strings.Join([]string{"ReportRuleRequest", string(data)}, " ")
}
