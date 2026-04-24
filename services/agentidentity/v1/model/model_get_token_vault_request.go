package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetTokenVaultRequest Request Object
type GetTokenVaultRequest struct {

	// The unique identifier of the token vault.
	TokenVaultId string `json:"token_vault_id"`
}

func (o GetTokenVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTokenVaultRequest struct{}"
	}

	return strings.Join([]string{"GetTokenVaultRequest", string(data)}, " ")
}
