package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVaultResponse Response Object
type CreateVaultResponse struct {
	Vault          *VaultCreateResource `json:"vault,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CreateVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultResponse struct{}"
	}

	return strings.Join([]string{"CreateVaultResponse", string(data)}, " ")
}
