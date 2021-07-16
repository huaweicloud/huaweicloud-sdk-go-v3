package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDiagnosisTaskDetailsRequest struct {
	// 诊断报告ID

	ReportId string `json:"report_id"`
}

func (o ShowDiagnosisTaskDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDiagnosisTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisTaskDetailsRequest", string(data)}, " ")
}
