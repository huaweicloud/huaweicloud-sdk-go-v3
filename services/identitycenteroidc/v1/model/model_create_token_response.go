package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTokenResponse Response Object
type CreateTokenResponse struct {
	TokenInfo      *TokenInfoDto `json:"token_info,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateTokenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTokenResponse struct{}"
	}

	return strings.Join([]string{"CreateTokenResponse", string(data)}, " ")
}
