package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainApiAclPolicyRequestBody
type UpdateDomainApiAclPolicyRequestBody struct {
	ApiAclPolicy *AclPolicyOption `json:"api_acl_policy"`
}

func (o UpdateDomainApiAclPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainApiAclPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainApiAclPolicyRequestBody", string(data)}, " ")
}
