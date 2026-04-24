package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceApiKeyResponse Response Object
type GetResourceApiKeyResponse struct {

	// API key associated with the requested resource
	ApiKey         *string `json:"api_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GetResourceApiKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceApiKeyResponse struct{}"
	}

	return strings.Join([]string{"GetResourceApiKeyResponse", string(data)}, " ")
}
