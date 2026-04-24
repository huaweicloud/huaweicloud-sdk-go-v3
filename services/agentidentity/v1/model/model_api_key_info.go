package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiKeyInfo struct {

	// The API key values. During rotation, any one of the keys can be used. If empty, a random string will be generated.
	ApiKey string `json:"api_key"`

	// The name of the API key.
	ApiKeyName *string `json:"api_key_name,omitempty"`
}

func (o ApiKeyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiKeyInfo struct{}"
	}

	return strings.Join([]string{"ApiKeyInfo", string(data)}, " ")
}
