package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportTaskInfoRequestBody struct {

	// **参数解释**: 导出的类型 **约束限制**: 不涉及 **取值范围**: - operational-report：月度运营报告  **默认取值**: 不涉及
	Type string `json:"type"`

	OperationalReportData *ExportTaskInfoRequestBodyOperationalReportData `json:"operational_report_data"`
}

func (o ExportTaskInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTaskInfoRequestBody struct{}"
	}

	return strings.Join([]string{"ExportTaskInfoRequestBody", string(data)}, " ")
}
