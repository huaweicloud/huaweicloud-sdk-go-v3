package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClientData struct {

	// 客户端语言。
	Language *string `json:"language,omitempty"`

	// 客户端版本。
	Version *string `json:"version,omitempty"`

	// 客户端ID。
	ClientId *string `json:"client_id,omitempty"`

	// 客户端地址。
	ClientAddr *string `json:"client_addr,omitempty"`

	// 订阅关系列表。
	Subscriptions *[]Subscription `json:"subscriptions,omitempty"`
}

func (o ClientData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientData struct{}"
	}

	return strings.Join([]string{"ClientData", string(data)}, " ")
}
