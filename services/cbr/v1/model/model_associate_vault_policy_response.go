package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateVaultPolicyResponse struct {
	AssociatePolicy *VaultPolicyResp `json:"associate_policy,omitempty" xml:"associate_policy"`
	HttpStatusCode  int              `json:"-"`
}

func (o AssociateVaultPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateVaultPolicyResponse", string(data)}, " ")
}
