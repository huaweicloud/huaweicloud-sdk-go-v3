package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisassociateVaultPolicyResponse struct {
	DissociatePolicy *VaultPolicyResp `json:"dissociate_policy,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o DisassociateVaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisassociateVaultPolicyResponse", string(data)}, " ")
}
