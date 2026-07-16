package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SfsTurboConnectionStatus 通过挂载网卡方式打通网络参数模型。
type SfsTurboConnectionStatus struct {

	// **参数解释**：sfsTurbo实例的名称。 **取值范围**：不涉及。
	Name string `json:"name"`

	// **参数解释**：sfsTurbo实例的ID。 **取值范围**：不涉及。
	SfsId string `json:"sfsId"`

	// **参数解释**：关联方式。 **取值范围**：可选值如下： - VpcPort：通过挂载网卡直通 - Peering：通过对等连接打通
	ConnectionType string `json:"connectionType"`

	// **参数解释**：SFS Turbo的访问地址。 **取值范围**：不涉及。
	IpAddr *string `json:"ipAddr,omitempty"`

	// **参数解释**：与SFS Turbo的连接状态信息。 **取值范围**：可选值如下： - Active：SFS连通状态正常 - Abnormal：SFS连通状态异常 - Creating：SFS连通状态创建关联中 - Deleting：SFS连通状态解除关联中
	Status *string `json:"status,omitempty"`
}

func (o SfsTurboConnectionStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SfsTurboConnectionStatus struct{}"
	}

	return strings.Join([]string{"SfsTurboConnectionStatus", string(data)}, " ")
}
