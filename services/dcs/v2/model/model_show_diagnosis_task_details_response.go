package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDiagnosisTaskDetailsResponse struct {
	// 诊断结果为异常的诊断项总数

	AbnormalItemSum *int32 `json:"abnormal_item_sum,omitempty"`
	// 诊断失败的诊断项总数

	FailedItemSum *int32 `json:"failed_item_sum,omitempty"`
	// 节点诊断报告列表

	DiagnosisNodeReportList *[]DiagnosisNodeReport `json:"diagnosis_node_report_list,omitempty"`
	HttpStatusCode          int                    `json:"-"`
}

func (o ShowDiagnosisTaskDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisTaskDetailsResponse", string(data)}, " ")
}
