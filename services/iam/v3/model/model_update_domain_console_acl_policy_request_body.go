package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateDomainConsoleAclPolicyRequestBody struct {
	ConsoleAclPolicy *AclPolicyOption `json:"console_acl_policy"`
}

func (o UpdateDomainConsoleAclPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainConsoleAclPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainConsoleAclPolicyRequestBody", string(data)}, " ")
}
