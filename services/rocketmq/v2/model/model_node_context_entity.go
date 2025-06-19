package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeContextEntity struct {

	// ID
	Id *string `json:"id,omitempty"`

	// broker名称
	BrokerName *string `json:"broker_name,omitempty"`

	// broker的ID
	BrokerId *int32 `json:"broker_id,omitempty"`

	// 地址
	Address *string `json:"address,omitempty"`

	// 公网地址
	PublicAddress *string `json:"public_address,omitempty"`
}

func (o NodeContextEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeContextEntity struct{}"
	}

	return strings.Join([]string{"NodeContextEntity", string(data)}, " ")
}
