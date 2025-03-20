package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttachedGroupPoliciesV5Response Response Object
type ListAttachedGroupPoliciesV5Response struct {

	// 身份策略列表。
	AttachedPolicies *[]AttachedPolicy `json:"attached_policies,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAttachedGroupPoliciesV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachedGroupPoliciesV5Response struct{}"
	}

	return strings.Join([]string{"ListAttachedGroupPoliciesV5Response", string(data)}, " ")
}
