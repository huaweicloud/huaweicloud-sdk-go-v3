package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelDetails struct {

	// **参数解释**： channel信息，包括客户端IP:Port到服务端IP:Port(channel_id)。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： channel数量。 **取值范围**： 不涉及。
	Number *int32 `json:"number,omitempty"`

	// **参数解释**： 消费者用户名，在开启ACL访问控制后返回真实用户名，未开启ACL时返回null。 **取值范围**： 不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**： connection信息，包括客户端IP:Port到服务端IP:Port。 **取值范围**： 不涉及。
	ConnectionName *string `json:"connection_name,omitempty"`

	// **参数解释**： 连接的消费者IP。 **取值范围**： 不涉及。
	PeerHost *string `json:"peer_host,omitempty"`

	// **参数解释**： 连接的消费者进程端口号。 **取值范围**： 不涉及。
	PeerPort *int32 `json:"peer_port,omitempty"`
}

func (o ChannelDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelDetails struct{}"
	}

	return strings.Join([]string{"ChannelDetails", string(data)}, " ")
}
