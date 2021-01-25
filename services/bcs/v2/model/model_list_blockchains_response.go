package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListBlockchainsResponse struct {
	// 服务实例简要信息
	Blockchains    *[]BlockchainInfo `json:"blockchains,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListBlockchainsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBlockchainsResponse struct{}"
	}

	return strings.Join([]string{"ListBlockchainsResponse", string(data)}, " ")
}
