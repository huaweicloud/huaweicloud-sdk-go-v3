package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttachedUserPoliciesV5Response Response Object
type ListAttachedUserPoliciesV5Response struct {

	// 身份策略列表。
	AttachedPolicies *[]AttachedPolicy `json:"attached_policies,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAttachedUserPoliciesV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachedUserPoliciesV5Response struct{}"
	}

	return strings.Join([]string{"ListAttachedUserPoliciesV5Response", string(data)}, " ")
}
