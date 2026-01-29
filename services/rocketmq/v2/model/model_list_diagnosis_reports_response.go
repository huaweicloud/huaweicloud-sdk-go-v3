package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDiagnosisReportsResponse Response Object
type ListDiagnosisReportsResponse struct {

	// **参数解释**： 诊断报告列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DiagnosisReportList *[]DiagnosisReportResp `json:"diagnosis_report_list,omitempty"`

	// **参数解释**： 报告数量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TotalNum       *int64 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDiagnosisReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisReportsResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnosisReportsResponse", string(data)}, " ")
}
