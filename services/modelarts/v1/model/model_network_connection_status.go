package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkConnectionStatus 当前网络与其他云服务的连接状态信息。
type NetworkConnectionStatus struct {

	// **参数解释**：Peer方式打通网络的状态信息列表。
	PeerConnectionStatus *[]PeerConnectionStatus `json:"peerConnectionStatus,omitempty"`

	// **参数解释**：网络可连通的SFS Turbo信息列表。
	SfsTurboStatus *[]SfsTurboConnectionStatus `json:"sfsTurboStatus,omitempty"`
}

func (o NetworkConnectionStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkConnectionStatus struct{}"
	}

	return strings.Join([]string{"NetworkConnectionStatus", string(data)}, " ")
}
