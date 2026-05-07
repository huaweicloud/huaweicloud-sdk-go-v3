package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityReportResponseInfo 查询报告总览页
type SecurityReportResponseInfo struct {

	// **参数解释**: 报告ID **取值范围**: 字符长度10-2147483647位
	ReportId *int32 `json:"report_id,omitempty"`

	// **参数解释**: 报告子ID **取值范围**: 字符长度10-2147483647位
	ReportSubId *int32 `json:"report_sub_id,omitempty"`

	// **参数解释**: 是否是默认的，默认的不能删除 **取值范围**: - true ：是。 - false ：否。
	DefaultReport *bool `json:"default_report,omitempty"`

	// **参数解释**： 最近生成时间，毫秒(如果返回值为null，代表暂未生成) **取值范围**: 不涉及
	LatestCreateTime *int64 `json:"latest_create_time,omitempty"`

	// **参数解释**: 报告名称 **取值范围**: 字符长度1-128位
	ReportName *string `json:"report_name,omitempty"`

	// **参数解释**: 报告类别 **取值范围**: - daily_report：安全日报 - weekly_report：安全周报 - monthly_report：安全月报 - custom_report：自定义报告
	ReportCategory *string `json:"report_category,omitempty"`

	// **参数解释**: 报告开启状态 **取值范围**:   - opened：开启   - closed：关闭
	ReportStatus *string `json:"report_status,omitempty"`

	// **参数解释**： 报告创建时间 **取值范围**: 不涉及
	ReportCreateTime *int64 `json:"report_create_time,omitempty"`

	// **参数解释**: 报告发送的时间段 **取值范围**:   - morning：代表0点到6点   - noon：代表6点到12点   - afternoon：代表12点到18点   - evening：代表18点到24点
	SendingPeriod *string `json:"sending_period,omitempty"`
}

func (o SecurityReportResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportResponseInfo struct{}"
	}

	return strings.Join([]string{"SecurityReportResponseInfo", string(data)}, " ")
}
