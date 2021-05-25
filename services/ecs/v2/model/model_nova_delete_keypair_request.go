package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NovaDeleteKeypairRequest struct {
	// 密钥名称。

	KeypairName string `json:"keypair_name"`
}

func (o NovaDeleteKeypairRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NovaDeleteKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaDeleteKeypairRequest", string(data)}, " ")
}
