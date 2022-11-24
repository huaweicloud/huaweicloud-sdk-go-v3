package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateMyActionTemplateResponse struct {

	// 三方算子ID
	TemplateName *string `json:"template_name,omitempty"`

	// 三方算子创建的时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 签名OBS地址，用于上传中英文算子包
	FuncPkgPath *string `json:"func_pkg_path,omitempty"`

	// 签名OBS地址，用于上传中英文算子logo
	FuncLogoPath map[string]string `json:"func_logo_path,omitempty"`

	// 签名OBS地址，用于上传中英文算子帮助文档
	FuncHelpPath map[string]string `json:"func_help_path,omitempty"`

	// 签名OBS地址，用于上传中英文算子测试报告
	FuncTestReportPath map[string]string `json:"func_test_report_path,omitempty"`

	// 签名OBS地址，用于上传中英文开源须知
	FuncOpensourceNoticePath map[string]string `json:"func_opensource_notice_path,omitempty"`

	// 签名OBS地址，用于上传中英文服务协议材料
	FuncSlaPath map[string]string `json:"func_sla_path,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMyActionTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMyActionTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateMyActionTemplateResponse", string(data)}, " ")
}
