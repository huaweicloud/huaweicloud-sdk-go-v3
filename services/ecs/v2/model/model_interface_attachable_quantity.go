package model

import (
	"encoding/json"

	"strings"
)

type InterfaceAttachableQuantity struct {
	// 云服务器剩余可挂载网卡数量

	FreeNic *int32 `json:"free_nic,omitempty"`
}

func (o InterfaceAttachableQuantity) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InterfaceAttachableQuantity struct{}"
	}

	return strings.Join([]string{"InterfaceAttachableQuantity", string(data)}, " ")
}
