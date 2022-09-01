package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSystemSecurityPoliciesResponse struct {

	// 系统安全策略列表。
	SystemSecurityPolicies *[]SystemSecurityPolicy `json:"system_security_policies,omitempty" xml:"system_security_policies"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSystemSecurityPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemSecurityPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListSystemSecurityPoliciesResponse", string(data)}, " ")
}
