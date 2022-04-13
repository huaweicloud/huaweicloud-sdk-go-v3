package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSystemSecurityPoliciesRequest struct {
}

func (o ListSystemSecurityPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemSecurityPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListSystemSecurityPoliciesRequest", string(data)}, " ")
}
