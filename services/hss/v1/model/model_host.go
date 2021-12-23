package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Host struct {
	// 云主机id

	AgentId *string `json:"agent_id,omitempty"`
	// 云主机id

	HostId *string `json:"host_id,omitempty"`
	// 云主机名称

	HostName *string `json:"host_name,omitempty"`
	// 云主机私有IP

	HostIp *string `json:"host_ip,omitempty"`
	// 云主机公网IP

	PublicIp *string `json:"public_ip,omitempty"`
	// 所属企业项目名称

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
	// 服务器组名称

	GroupName *string `json:"group_name,omitempty"`
	// 服务到期时间

	ExpireTime *int64 `json:"expire_time,omitempty"`
	// 策略组名称

	PolicyGroupName *string `json:"policy_group_name,omitempty"`
	// Agent状态，包含如下4种。   - ACTIVE ：正在运行。   - SHUTOFF ：关机。   - BUILDING ：创建中。   - ERROR ：故障。

	HostStatus *string `json:"host_status,omitempty"`
	// Agent状态，包含如下3种。   - uninstall ：未注册。   - online ：在线。   - offline ：离线。

	AgentStatus *string `json:"agent_status,omitempty"`
	// 主机开通的版本，包含如下5种输入。   - hss.version.null ：无。   - hss.version.basic ：基础版。   - hss.version.enterprise ：企业版。   - hss.version.premium ：旗舰版。   - hss.version.wtp ：网页防篡改版。

	Version *string `json:"version,omitempty"`
	// 防护状态，包含如下2种。 - closed ：关闭。 - opened ：开启。

	ProtectStatus *string `json:"protect_status,omitempty"`
	// 系统镜像

	OsImage *string `json:"os_image,omitempty"`
	// 系统类型

	OsType *string `json:"os_type,omitempty"`
	// 操作系统位数

	OsBit *string `json:"os_bit,omitempty"`
	// 云主机安全检测结果，包含如下3种。 - undetect ：未检测。 - clean ：无风险。 - risk ：有风险。

	DetectResult *string `json:"detect_result,omitempty"`
	// 资产风险个数

	RiskPortNum *int32 `json:"risk_port_num,omitempty"`
	// 漏洞风险个数

	RiskVulNum *int32 `json:"risk_vul_num,omitempty"`
	// 入侵风险个数

	RiskIntrusionNum *int32 `json:"risk_intrusion_num,omitempty"`
	// 基线风险个数

	RiskBaselineNum *int32 `json:"risk_baseline_num,omitempty"`
	// 收费模式，包含如下2种。   - packet_cycle ：包年/包月。   - on_demand ：按需。

	ChargingMode *string `json:"charging_mode,omitempty"`
	// 云服务资源实例ID（UUID）

	ResourceId *string `json:"resource_id,omitempty"`
}

func (o Host) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Host struct{}"
	}

	return strings.Join([]string{"Host", string(data)}, " ")
}
