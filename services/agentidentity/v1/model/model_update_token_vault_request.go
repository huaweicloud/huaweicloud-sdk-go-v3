package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTokenVaultRequest Request Object
type UpdateTokenVaultRequest struct {

	// The unique identifier of the token vault.
	TokenVaultId string `json:"token_vault_id"`

	Body *UpdateTokenVaultReqBody `json:"body,omitempty"`
}

func (o UpdateTokenVaultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTokenVaultRequest struct{}"
	}

	return strings.Join([]string{"UpdateTokenVaultRequest", string(data)}, " ")
}
