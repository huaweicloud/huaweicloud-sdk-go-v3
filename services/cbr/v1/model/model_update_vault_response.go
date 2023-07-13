package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVaultResponse Response Object
type UpdateVaultResponse struct {
	Vault          *Vault `json:"vault,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVaultResponse struct{}"
	}

	return strings.Join([]string{"UpdateVaultResponse", string(data)}, " ")
}
