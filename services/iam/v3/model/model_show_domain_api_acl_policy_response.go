package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDomainApiAclPolicyResponse struct {
	ApiAclPolicy   *AclPolicyResult `json:"api_acl_policy,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDomainApiAclPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDomainApiAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainApiAclPolicyResponse", string(data)}, " ")
}
