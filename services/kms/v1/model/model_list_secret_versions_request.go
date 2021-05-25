package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSecretVersionsRequest struct {
	// 凭据的资源标识符。

	SecretId string `json:"secret_id"`
}

func (o ListSecretVersionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSecretVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretVersionsRequest", string(data)}, " ")
}
