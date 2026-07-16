package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkConnection 当前网络与其他网络的连接信息。
type NetworkConnection struct {

	// **参数解释**：Peer方式打通网络列表。
	PeerConnectionList *[]PeerConnectionItem `json:"peerConnectionList,omitempty"`
}

func (o NetworkConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkConnection struct{}"
	}

	return strings.Join([]string{"NetworkConnection", string(data)}, " ")
}
