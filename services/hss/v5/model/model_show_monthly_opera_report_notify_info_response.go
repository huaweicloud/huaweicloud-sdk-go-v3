package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMonthlyOperaReportNotifyInfoResponse Response Object
type ShowMonthlyOperaReportNotifyInfoResponse struct {

	// close:不显示弹框 open：显示弹框
	Status *string `json:"status,omitempty"`

	// 称号 -vul-fix-master: 补洞大师 -vul-fix-expert: 漏洞修复小能手 -baseline-fix: 风险配置修复高手 -malware-file: 防病毒先锋 -ransomware-event: 防勒索达人 -web-tamper-event: 网站守卫
	Title *string `json:"title,omitempty"`

	// 当前用户唯一标识，报告时间，返回字符串：2024-04
	ReportId *string `json:"report_id,omitempty"`

	// 当前月份，6
	CurrentMonth   *int32 `json:"current_month,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMonthlyOperaReportNotifyInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonthlyOperaReportNotifyInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowMonthlyOperaReportNotifyInfoResponse", string(data)}, " ")
}
