package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateExchangeBody struct {

	// Exchange名称
	Name string `json:"name"`

	// 类型（direct、fanout、topic、headers）
	Type string `json:"type"`

	// 是否持久化[（AMQP版本默认持久化，不涉及此参数）](tag:hws,hws_hk)。
	Durable *bool `json:"durable,omitempty"`

	// 是否自动删除
	AutoDelete bool `json:"auto_delete"`

	// 内部Exchange[（AMQP版本不支持内部Exchange，不涉及此参数）](tag:hws,hws_hk)。
	Internal *bool `json:"internal,omitempty"`
}

func (o CreateExchangeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExchangeBody struct{}"
	}

	return strings.Join([]string{"CreateExchangeBody", string(data)}, " ")
}
