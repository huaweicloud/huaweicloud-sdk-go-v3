package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDiagnosisReportResponse Response Object
type BatchDeleteDiagnosisReportResponse struct {

	// **参数解释**： 诊断报告ID列表。 **取值范围**： 不涉及。
	ReportIdList   *[]string `json:"report_id_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteDiagnosisReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDiagnosisReportResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteDiagnosisReportResponse", string(data)}, " ")
}
