package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTestReportCustomDetailRequest Request Object
type SaveTestReportCustomDetailRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 计划uri
	VersionUri string `json:"version_uri"`

	// 测试报告Uri
	ReportUri string `json:"report_uri"`

	Body *TestReportCustomDetailInfo `json:"body,omitempty"`
}

func (o SaveTestReportCustomDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTestReportCustomDetailRequest struct{}"
	}

	return strings.Join([]string{"SaveTestReportCustomDetailRequest", string(data)}, " ")
}
