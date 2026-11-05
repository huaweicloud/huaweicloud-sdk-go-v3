package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupabaseApplicationApiKeyValueResponse Response Object
type ShowSupabaseApplicationApiKeyValueResponse struct {

	// ApiKey明文。
	ApiKey         *string `json:"api_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSupabaseApplicationApiKeyValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupabaseApplicationApiKeyValueResponse struct{}"
	}

	return strings.Join([]string{"ShowSupabaseApplicationApiKeyValueResponse", string(data)}, " ")
}
