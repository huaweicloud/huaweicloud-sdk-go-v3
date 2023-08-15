package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateVaultPolicyResponse Response Object
type AssociateVaultPolicyResponse struct {
	AssociatePolicy *VaultPolicyResp `json:"associate_policy,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o AssociateVaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateVaultPolicyResponse", string(data)}, " ")
}
