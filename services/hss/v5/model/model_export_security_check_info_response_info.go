package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSecurityCheckInfoResponseInfo 导出的配置检测信息
type ExportSecurityCheckInfoResponseInfo struct {

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器公网IP
	HostPublicIp *string `json:"host_public_ip,omitempty"`

	// 服务器私网IP
	HostPrivateIp *string `json:"host_private_ip,omitempty"`

	// 配置检查（基线）的类型,Linux系统支持的基线一般check_type和check_name相同,例如SSH、CentOS 7。 Windows系统支持的基线一般check_type和check_name不相同，例如check_name为Windows的配置检查（基线），它的check_type包含Windows Server 2019 R2、Windows Server 2016 R2等。
	CheckType *string `json:"check_type,omitempty"`

	// 配置检查（基线）的名称，例如SSH、CentOS 7、Windows
	CheckName *string `json:"check_name,omitempty"`

	// 基线标准 - cn_standard#等保合规标准 - hw_standard#云安全实践标准
	Standard *string `json:"standard,omitempty"`

	// 检查项（检查规则）名称
	CheckRuleName *string `json:"check_rule_name,omitempty"`

	// 配置检查（基线）的路径信息
	ExecutableFilePath *string `json:"executable_file_path,omitempty"`

	// 风险等级，包含如下:   - Low : 低危   - Medium : 中危   - High : 高危
	Severity *string `json:"severity,omitempty"`

	// 检测结果  - pass - failed
	ScanResult *string `json:"scan_result,omitempty"`

	// 状态  - safe : 无需处理 - ignored : 已忽略 - unhandled : 未处理
	Status *string `json:"status,omitempty"`

	// 规则描述
	CheckRuleDesc *string `json:"check_rule_desc,omitempty"`

	// 审计描述
	Audit *string `json:"audit,omitempty"`

	// 修改建议
	Remediation *string `json:"remediation,omitempty"`

	// 检测用例信息
	CheckInfoList *string `json:"check_info_list,omitempty"`

	// 首次扫描时间，时间戳单位：毫秒
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o ExportSecurityCheckInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSecurityCheckInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"ExportSecurityCheckInfoResponseInfo", string(data)}, " ")
}
