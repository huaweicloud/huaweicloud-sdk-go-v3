package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoliciesV5Response Response Object
type ListPoliciesV5Response struct {

	// 身份策略列表。
	Policies *[]Policy `json:"policies,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPoliciesV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesV5Response struct{}"
	}

	return strings.Join([]string{"ListPoliciesV5Response", string(data)}, " ")
}
