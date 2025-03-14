package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainConsoleAclPolicyResponse Response Object
type UpdateDomainConsoleAclPolicyResponse struct {
	ConsoleAclPolicy *ConsoleAclPolicyResult `json:"console_acl_policy,omitempty"`
	HttpStatusCode   int                     `json:"-"`
}

func (o UpdateDomainConsoleAclPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainConsoleAclPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainConsoleAclPolicyResponse", string(data)}, " ")
}
