package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PeerConnectionStatus Peer方式打通网络状态参数模型。
type PeerConnectionStatus struct {

	// **参数解释**：对端的VPC ID。 **取值范围**：不涉及。
	PeerVpcId string `json:"peerVpcId"`

	// **参数解释**：对端的子网ID。 **取值范围**：不涉及。
	PeerSubnetId string `json:"peerSubnetId"`

	// **参数解释**：创建默认路由的开关。 **取值范围**：默认为false不创建。
	DefaultGateWay *bool `json:"defaultGateWay,omitempty"`

	// **参数解释**：网络的连接状态。 **取值范围**：可选值如下： - Connecting：网络连接中 - Active：网络连接正常 - Abnormal：网络连接不正常
	Phase *string `json:"phase,omitempty"`
}

func (o PeerConnectionStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerConnectionStatus struct{}"
	}

	return strings.Join([]string{"PeerConnectionStatus", string(data)}, " ")
}
