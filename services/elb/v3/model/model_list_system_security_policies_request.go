package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemSecurityPoliciesRequest Request Object
type ListSystemSecurityPoliciesRequest struct {
}

func (o ListSystemSecurityPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemSecurityPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListSystemSecurityPoliciesRequest", string(data)}, " ")
}
