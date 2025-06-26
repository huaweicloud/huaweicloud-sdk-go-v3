package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BgpPeer struct {

	// 对端IP地址
	PeerIpAddress *string `json:"peer_ip_address,omitempty"`

	// 对端AS号
	PeerAsn *int64 `json:"peer_asn,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 状态周期
	StateDuration *string `json:"state_duration,omitempty"`

	// 接收到的路由数量
	NumReceivedRoutes *int32 `json:"num_received_routes,omitempty"`

	// 接收到的消息数量
	NumMessageReceived *int32 `json:"num_message_received,omitempty"`

	// 已发送的消息数量
	NumMessageSent *int32 `json:"num_message_sent,omitempty"`
}

func (o BgpPeer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BgpPeer struct{}"
	}

	return strings.Join([]string{"BgpPeer", string(data)}, " ")
}
