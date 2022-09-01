package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSecurityPoliciesResponse struct {

	// 自定义安全策略列表返回对象。
	SecurityPolicies *[]SecurityPolicy `json:"security_policies,omitempty" xml:"security_policies"`

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecurityPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityPoliciesResponse", string(data)}, " ")
}
