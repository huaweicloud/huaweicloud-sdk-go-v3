package model

import (
	"encoding/json"

	"strings"
)

type BlockchainInfo struct {
	// 服务实例ID

	Id *string `json:"id,omitempty"`
	// 服务实例名

	Name *string `json:"name,omitempty"`
}

func (o BlockchainInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BlockchainInfo struct{}"
	}

	return strings.Join([]string{"BlockchainInfo", string(data)}, " ")
}
