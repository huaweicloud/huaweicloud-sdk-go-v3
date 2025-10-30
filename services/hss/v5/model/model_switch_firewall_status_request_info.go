package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchFirewallStatusRequestInfo struct {

	// **参数解释**： 防火墙开启状态 **约束限制**： 不涉及 **取值范围**：   - Opened：开启windows防火墙   - Closed：关闭windows防火墙 **默认取值**： 不涉及
	FirewallStatus string `json:"firewall_status"`

	// **参数解释**： 主机ID列表 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	HostIdList []string `json:"host_id_list"`
}

func (o SwitchFirewallStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchFirewallStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"SwitchFirewallStatusRequestInfo", string(data)}, " ")
}
