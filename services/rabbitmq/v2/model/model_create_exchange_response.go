package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateExchangeResponse Response Object
type CreateExchangeResponse struct {

	// 是否持久化
	Durable *bool `json:"durable,omitempty"`

	// 是否是默认Exchange
	Default *bool `json:"default,omitempty"`

	// 是否是内部Exchange
	Internal *bool `json:"internal,omitempty"`

	// Exchange名称
	Name *string `json:"name,omitempty"`

	// 是否自动删除
	AutoDelete *bool `json:"auto_delete,omitempty"`

	// Exchange类型
	Type *string `json:"type,omitempty"`

	// 所属Vhost
	Vhost          *string `json:"vhost,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateExchangeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExchangeResponse struct{}"
	}

	return strings.Join([]string{"CreateExchangeResponse", string(data)}, " ")
}
