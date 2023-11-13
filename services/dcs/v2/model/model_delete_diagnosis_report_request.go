package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDiagnosisReportRequest struct {

	// 诊断报告ID
	ReportIdList *[]string `json:"report_id_list,omitempty"`
}

func (o DeleteDiagnosisReportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDiagnosisReportRequest struct{}"
	}

	return strings.Join([]string{"DeleteDiagnosisReportRequest", string(data)}, " ")
}
