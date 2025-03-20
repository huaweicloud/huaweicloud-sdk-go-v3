package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAttachedAgencyPoliciesV5Response Response Object
type ListAttachedAgencyPoliciesV5Response struct {

	// 身份策略列表。
	AttachedPolicies *[]AttachedPolicy `json:"attached_policies,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAttachedAgencyPoliciesV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAttachedAgencyPoliciesV5Response struct{}"
	}

	return strings.Join([]string{"ListAttachedAgencyPoliciesV5Response", string(data)}, " ")
}
