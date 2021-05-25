package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateDomainConsoleAclPolicyRequest struct {
	// 账号ID，获取方式请参见：[获取账号ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	DomainId string `json:"domain_id"`

	Body *UpdateDomainConsoleAclPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainConsoleAclPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateDomainConsoleAclPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainConsoleAclPolicyRequest", string(data)}, " ")
}
