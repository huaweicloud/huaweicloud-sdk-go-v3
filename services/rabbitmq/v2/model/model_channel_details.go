package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelDetails struct {

	// channel信息，包括客户端IP:Port到服务端IP:Port(channel_id)。
	Name *string `json:"name,omitempty"`

	// channel数量
	Number *int32 `json:"number,omitempty"`

	// 消费者用户名，在开启ACL访问控制后返回真实用户名，未开启ACL时返回null。
	User *string `json:"user,omitempty"`

	// connection信息，包括客户端IP:Port到服务端IP:Port。
	ConnectionName *string `json:"connection_name,omitempty"`

	// 连接的消费者IP
	PeerHost *string `json:"peer_host,omitempty"`

	// 连接的消费者进程端口号
	PeerPort *int32 `json:"peer_port,omitempty"`
}

func (o ChannelDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelDetails struct{}"
	}

	return strings.Join([]string{"ChannelDetails", string(data)}, " ")
}
