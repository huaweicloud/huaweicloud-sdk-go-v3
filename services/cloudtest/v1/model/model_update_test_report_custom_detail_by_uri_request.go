package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestReportCustomDetailByUriRequest Request Object
type UpdateTestReportCustomDetailByUriRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 计划uri
	VersionUri string `json:"version_uri"`

	// 测试报告Uri
	ReportUri string `json:"report_uri"`

	// 测试报告自定义模块Uri
	CustomInfoUri string `json:"custom_info_uri"`

	Body *TestReportCustomDetailInfo `json:"body,omitempty"`
}

func (o UpdateTestReportCustomDetailByUriRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestReportCustomDetailByUriRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestReportCustomDetailByUriRequest", string(data)}, " ")
}
