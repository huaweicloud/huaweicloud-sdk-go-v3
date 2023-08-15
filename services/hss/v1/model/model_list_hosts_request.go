package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsRequest Request Object
type ListHostsRequest struct {

	// 主机开通的版本，包含如下5种输入。   - hss.version.null ：无。   - hss.version.basic ：基础版。   - hss.version.enterprise ：企业版。   - hss.version.premium ：旗舰版。   - hss.version.wtp ：网页防篡改版。
	Version *string `json:"version,omitempty"`

	// Agent状态，包含如下3种。   - not_register ：未注册。   - online ：在线。   - offline ：离线。
	AgentStatus *string `json:"agent_status,omitempty"`

	// Agent状态，包含如下4种。   - ACTIVE ：正在运行。   - SHUTOFF ：关机。   - BUILDING ：创建中。   - ERROR ：故障。
	HostStatus *string `json:"host_status,omitempty"`

	// 防护状态，包含如下2种。   - closed ：关闭。   - opened ：开启。
	ProtectStatus *string `json:"protect_status,omitempty"`

	// 防护状态，包含如下3种。   - undetect ：未检测。   - clean ：无风险。   - risk ：有风险。
	DetectResult *string `json:"detect_result,omitempty"`

	// 云主机名称
	HostName *string `json:"host_name,omitempty"`

	// 云主机私有IP
	HostIp *string `json:"host_ip,omitempty"`

	// 云主机公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// 操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// 收费模式，包含如下2种。   - packet_cycle ：包年/包月。   - on_demand ：按需。
	ChargingMode *string `json:"charging_mode,omitempty"`

	// 默认10
	Limit *int32 `json:"limit,omitempty"`

	// 默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListHostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsRequest struct{}"
	}

	return strings.Join([]string{"ListHostsRequest", string(data)}, " ")
}
