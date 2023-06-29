package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiagnosisTaskDetailsRequest Request Object
type ShowDiagnosisTaskDetailsRequest struct {

	// 诊断报告ID
	ReportId string `json:"report_id"`
}

func (o ShowDiagnosisTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiagnosisTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowDiagnosisTaskDetailsRequest", string(data)}, " ")
}
