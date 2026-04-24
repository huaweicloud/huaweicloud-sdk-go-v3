package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetTokenVaultResponse Response Object
type GetTokenVaultResponse struct {
	TokenVault     *TokenVault `json:"token_vault,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o GetTokenVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTokenVaultResponse struct{}"
	}

	return strings.Join([]string{"GetTokenVaultResponse", string(data)}, " ")
}
