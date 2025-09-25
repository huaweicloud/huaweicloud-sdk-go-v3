package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostRiskResponseInfo struct {

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// Agent状态，包含如下5种。   - not_installed ：未安装。   - online ：在线。   - offline ：离线。   - install_failed ：安装失败。   - installing ：安装中。
	AgentStatus *string `json:"agent_status,omitempty"`

	// 安装结果，包含如下12种。   - install_succeed ：安装成功。   - network_access_timeout ：网络不通，访问超时。   - invalid_port ：无效端口。   - auth_failed ：认证错误，口令不正确。   - permission_denied ：权限错误，被拒绝。   - no_available_vpc ：没有相同VPC的agent在线虚拟机。   - install_exception ：安装异常。   - invalid_param ：参数错误。   - install_failed ：安装失败。   - package_unavailable ：安装包失效。   - os_type_not_support ：系统类型错误。   - os_arch_not_support ：架构类型错误。
	InstallResultCode *string `json:"install_result_code,omitempty"`

	// 主机开通的版本，包含如下7种输入。   - hss.version.null ：无。   - hss.version.basic ：基础版。   - hss.version.advanced ：专业版。   - hss.version.enterprise ：企业版。   - hss.version.premium ：旗舰版。   - hss.version.wtp ：网页防篡改版。   - hss.version.container.enterprise ：容器版。
	Version *string `json:"version,omitempty"`

	// 防护状态，包含如下2种。 - closed ：未防护。 - opened ：防护中。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 云主机安全检测结果，包含如下4种。 - undetected ：未检测。 - clean ：无风险。 - risk ：有风险。 - scanning ：检测中。
	DetectResult *string `json:"detect_result,omitempty"`

	// 资产风险
	Asset *int32 `json:"asset,omitempty"`

	// 漏洞风险
	Vulnerability *int32 `json:"vulnerability,omitempty"`

	// 基线风险
	Baseline *int32 `json:"baseline,omitempty"`

	// 入侵风险
	Intrusion *int32 `json:"intrusion,omitempty"`
}

func (o HostRiskResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostRiskResponseInfo struct{}"
	}

	return strings.Join([]string{"HostRiskResponseInfo", string(data)}, " ")
}
