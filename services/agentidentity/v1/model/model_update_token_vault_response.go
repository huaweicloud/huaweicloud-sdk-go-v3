package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTokenVaultResponse Response Object
type UpdateTokenVaultResponse struct {
	TokenVault     *TokenVault `json:"token_vault,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateTokenVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTokenVaultResponse struct{}"
	}

	return strings.Join([]string{"UpdateTokenVaultResponse", string(data)}, " ")
}
