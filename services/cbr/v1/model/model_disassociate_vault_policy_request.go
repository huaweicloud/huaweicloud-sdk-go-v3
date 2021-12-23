package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateVaultPolicyRequest struct {
	// 存储库ID

	VaultId string `json:"vault_id"`

	Body *VaultDissociate `json:"body,omitempty"`
}

func (o DisassociateVaultPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVaultPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisassociateVaultPolicyRequest", string(data)}, " ")
}
