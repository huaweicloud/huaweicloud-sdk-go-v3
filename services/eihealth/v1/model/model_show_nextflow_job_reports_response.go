package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowNextflowJobReportsResponse struct {

	// 作业报告文件列表
	ReportFiles    *[]NextflowJobReportFile `json:"report_files,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowNextflowJobReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNextflowJobReportsResponse struct{}"
	}

	return strings.Join([]string{"ShowNextflowJobReportsResponse", string(data)}, " ")
}
