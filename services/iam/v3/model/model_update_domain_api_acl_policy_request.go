package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainApiAclPolicyRequest Request Object
type UpdateDomainApiAclPolicyRequest struct {

	// 账号ID，获取方式请参见：[获取账号ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。
	DomainId string `json:"domain_id"`

	Body *UpdateDomainApiAclPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateDomainApiAclPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainApiAclPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainApiAclPolicyRequest", string(data)}, " ")
}
