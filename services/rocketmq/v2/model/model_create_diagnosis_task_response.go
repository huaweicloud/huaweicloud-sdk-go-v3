package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDiagnosisTaskResponse Response Object
type CreateDiagnosisTaskResponse struct {

	// **参数解释**： 报告ID。 **取值范围**： 不涉及。
	ReportId       *string `json:"report_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDiagnosisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskResponse", string(data)}, " ")
}
