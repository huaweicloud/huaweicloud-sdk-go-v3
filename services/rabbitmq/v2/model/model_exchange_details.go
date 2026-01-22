package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExchangeDetails struct {

	// 是否持久化
	Durable *bool `json:"durable,omitempty"`

	// 是否是默认Exchange
	Default *bool `json:"default,omitempty"`

	// 是否是内部Exchange
	Internal *bool `json:"internal,omitempty"`

	// 参数列表
	Arguments *interface{} `json:"arguments,omitempty"`

	// Exchange名称
	Name *string `json:"name,omitempty"`

	// 是否自动删除
	AutoDelete *bool `json:"auto_delete,omitempty"`

	// **参数解释**： Exchange类型。 **取值范围**： - direct：该类型Exchange会将消息路由到Routing Key完全匹配的Queue中。 - fanout：该类型Exchange会将消息路由到所有与其绑定的Queue中。 - topic：该类型Exchange将Routing Key进行通配符匹配，然后将消息路由到匹配成功的Queue中。 - headers：该类型Exchange与Routing Key无关，而与消息中的Headers属性信息相关。Exchange根据消息中的Headers属性键值对和绑定的属性键值对进行匹配，根据匹配情况路由消息。
	Type *string `json:"type,omitempty"`

	// 所属Vhost
	Vhost *string `json:"vhost,omitempty"`
}

func (o ExchangeDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExchangeDetails struct{}"
	}

	return strings.Join([]string{"ExchangeDetails", string(data)}, " ")
}
