package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityResourcePermissionPoliciesResponse Response Object
type BatchDeleteSecurityResourcePermissionPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecurityResourcePermissionPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityResourcePermissionPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityResourcePermissionPoliciesResponse", string(data)}, " ")
}
