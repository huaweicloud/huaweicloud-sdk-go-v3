package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVaultResponse struct {
	Vault          *Vault `json:"vault,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateVaultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVaultResponse struct{}"
	}

	return strings.Join([]string{"CreateVaultResponse", string(data)}, " ")
}
