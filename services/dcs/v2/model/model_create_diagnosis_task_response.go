package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDiagnosisTaskResponse struct {

	// 报告ID
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
