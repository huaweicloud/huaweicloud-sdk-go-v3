package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PeerConnectionItem Peer方式打通网络参数模型。
type PeerConnectionItem struct {

	// **参数解释**：对端的VPC ID。 **取值范围**：不涉及。
	PeerVpcId string `json:"peerVpcId"`

	// **参数解释**：对端的子网ID。 **取值范围**：不涉及。
	PeerSubnetId string `json:"peerSubnetId"`

	// **参数解释**：创建默认路由的开关，默认为false不创建。 **取值范围**：不涉及。
	DefaultGateWay *bool `json:"defaultGateWay,omitempty"`
}

func (o PeerConnectionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeerConnectionItem struct{}"
	}

	return strings.Join([]string{"PeerConnectionItem", string(data)}, " ")
}
