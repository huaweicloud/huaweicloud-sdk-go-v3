package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateVaultPolicyRequest struct {

	// 存储库ID
	VaultId string `json:"vault_id" xml:"vault_id"`

	Body *VaultDissociate `json:"body,omitempty" xml:"body"`
}

func (o DisassociateVaultPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVaultPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisassociateVaultPolicyRequest", string(data)}, " ")
}
