package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebTamperEventHostInfo **参数解释**: 网页防篡改事件信息对应的主机信息 **取值范围**: 不涉及
type WebTamperEventHostInfo struct {

	// **参数解释**: 主机id **取值范围**: 字符长度1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 主机名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机公网ip **取值范围**: 字符长度1-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**: 主机私网ip **取值范围**: 字符长度1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**: 主机的资产重要性 **取值范围**: - important：重要资产 - common：一般资产 - test：测试资产
	AssetValue *string `json:"asset_value,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： 包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**: 系统名称 **取值范围**: 字符长度0-128位
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**: 服务器状态 **取值范围**: 包含如下4种。 - ACTIVE ：运行中。 - SHUTOFF ：关机。 - BUILDING ：创建中。 - ERROR ：故障。
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**: Agent状态 **取值范围**: 包含如下6种。 - installed ：已安装。 - not_installed ：未安装。 - online ：在线。 - offline ：离线。 - install_failed ：安装失败。 - installing ：安装中。
	AgentStatus *string `json:"agent_status,omitempty"`
}

func (o WebTamperEventHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperEventHostInfo struct{}"
	}

	return strings.Join([]string{"WebTamperEventHostInfo", string(data)}, " ")
}
