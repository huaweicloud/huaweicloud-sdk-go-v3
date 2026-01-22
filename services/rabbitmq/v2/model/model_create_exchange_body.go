package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateExchangeBody struct {

	// Exchange名称
	Name string `json:"name"`

	// **参数解释**： Exchange类型。 **约束限制**： 不涉及。 **取值范围**： - direct：该类型Exchange会将消息路由到Routing Key完全匹配的Queue中。 - fanout：该类型Exchange会将消息路由到所有与其绑定的Queue中。 - topic：该类型Exchange将Routing Key进行通配符匹配，然后将消息路由到匹配成功的Queue中。 - headers：该类型Exchange与Routing Key无关，而与消息中的Headers属性信息相关。Exchange根据消息中的Headers属性键值对和绑定的属性键值对进行匹配，根据匹配情况路由消息。 **默认取值**： 不涉及。
	Type string `json:"type"`

	// 是否持久化[（AMQP版本默认持久化，不涉及此参数）](tag:hws,hws_hk,hws_eu)。
	Durable *bool `json:"durable,omitempty"`

	// 是否自动删除
	AutoDelete bool `json:"auto_delete"`

	// 内部Exchange[（AMQP版本不支持内部Exchange，不涉及此参数）](tag:hws,hws_hk,hws_eu)。
	Internal *bool `json:"internal,omitempty"`

	// 参数列表
	Arguments *interface{} `json:"arguments,omitempty"`
}

func (o CreateExchangeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExchangeBody struct{}"
	}

	return strings.Join([]string{"CreateExchangeBody", string(data)}, " ")
}
