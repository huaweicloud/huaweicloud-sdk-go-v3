package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DisassociateVaultPolicyResponse struct {
	DissociatePolicy *VaultPolicyResp `json:"dissociate_policy,omitempty"`
	HttpStatusCode   int              `json:"-"`
}

func (o DisassociateVaultPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateVaultPolicyResponse struct{}"
	}

	return strings.Join([]string{"DisassociateVaultPolicyResponse", string(data)}, " ")
}
