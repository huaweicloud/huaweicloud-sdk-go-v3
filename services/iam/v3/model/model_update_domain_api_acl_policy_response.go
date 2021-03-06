package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateDomainApiAclPolicyResponse struct {
	ApiAclPolicy   *AclPolicyResult `json:"api_acl_policy,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateDomainApiAclPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainApiAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainApiAclPolicyResponse", string(data)}, " ")
}
