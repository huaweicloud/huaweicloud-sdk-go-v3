package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListKeypairsResponse struct {
	// SSH密钥对信息详情

	Keypairs       *[]Keypairs `json:"keypairs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListKeypairsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListKeypairsResponse struct{}"
	}

	return strings.Join([]string{"ListKeypairsResponse", string(data)}, " ")
}
