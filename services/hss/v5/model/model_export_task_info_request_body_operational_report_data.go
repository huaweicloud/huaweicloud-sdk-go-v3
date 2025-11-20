package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTaskInfoRequestBodyOperationalReportData **参数解释**: 导出的报告参数数据
type ExportTaskInfoRequestBodyOperationalReportData struct {

	// **参数解释**: 导出的报告的id **约束限制**: 不涉及 **取值范围**: 字符长度1-32位 **默认取值**: 不涉及
	ReportId string `json:"report_id"`
}

func (o ExportTaskInfoRequestBodyOperationalReportData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTaskInfoRequestBodyOperationalReportData struct{}"
	}

	return strings.Join([]string{"ExportTaskInfoRequestBodyOperationalReportData", string(data)}, " ")
}
