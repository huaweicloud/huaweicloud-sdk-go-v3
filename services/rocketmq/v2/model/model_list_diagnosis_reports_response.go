package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnosisReportsResponse Response Object
type ListDiagnosisReportsResponse struct {

	// **参数解释**： 诊断报告列表。 **取值范围**： 不涉及。
	DiagnosisReportList *[]DiagnosisReportResp `json:"diagnosis_report_list,omitempty"`
	HttpStatusCode      int                    `json:"-"`
}

func (o ListDiagnosisReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisReportsResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnosisReportsResponse", string(data)}, " ")
}
