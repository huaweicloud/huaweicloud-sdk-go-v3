package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowImageSecurityReportStatisticRequest Request Object
type ShowImageSecurityReportStatisticRequest struct {

	// Region ID
	Region string `json:"region"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 安全报告类型，包含如下:   - vulnerability: 漏洞信息   - risk-config: 基线配置信息   - malwares: 恶意文件信息   - apps: 软件信息   - files: 文件信息   - sensitive-info: 敏感信息   - non-compliant-app: 软件合规信息   - build-command-risks: 镜像构建指令风险
	SecurityReportType string `json:"security_report_type"`

	Body *BatchExportSwrBaselineInfoRequestInfo `json:"body,omitempty"`
}

func (o ShowImageSecurityReportStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageSecurityReportStatisticRequest struct{}"
	}

	return strings.Join([]string{"ShowImageSecurityReportStatisticRequest", string(data)}, " ")
}
