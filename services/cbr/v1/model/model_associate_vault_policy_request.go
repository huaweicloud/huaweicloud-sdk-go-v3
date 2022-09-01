package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AssociateVaultPolicyRequest struct {

	// 存储库ID
	VaultId string `json:"vault_id" xml:"vault_id"`

	Body *VaultAssociate `json:"body,omitempty" xml:"body"`
}

func (o AssociateVaultPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateVaultPolicyRequest struct{}"
	}

	return strings.Join([]string{"AssociateVaultPolicyRequest", string(data)}, " ")
}
