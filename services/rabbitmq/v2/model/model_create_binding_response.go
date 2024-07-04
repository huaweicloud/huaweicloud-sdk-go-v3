package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBindingResponse Response Object
type CreateBindingResponse struct {

	// 绑定对象
	Source *string `json:"source,omitempty"`

	// 绑定Exchange或者Queue
	DestinationType *string `json:"destination_type,omitempty"`

	// 要投递的Exchange或Queue名称
	Destination *string `json:"destination,omitempty"`

	// 绑定键值，用于告知Exchange应该将消息投递到哪些Queue中
	RoutingKey     *string `json:"routing_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBindingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBindingResponse struct{}"
	}

	return strings.Join([]string{"CreateBindingResponse", string(data)}, " ")
}
