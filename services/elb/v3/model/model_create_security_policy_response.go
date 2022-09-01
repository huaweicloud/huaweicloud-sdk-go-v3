package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateSecurityPolicyResponse struct {
	SecurityPolicy *SecurityPolicy `json:"security_policy,omitempty" xml:"security_policy"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSecurityPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyResponse", string(data)}, " ")
}
