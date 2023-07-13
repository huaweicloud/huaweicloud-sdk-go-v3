package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedPoliciesInPermissionSetResponse Response Object
type ListManagedPoliciesInPermissionSetResponse struct {

	// AttachedManagedPolicy数据类型对象的数组
	AttachedManagedPolicies *[]AttachedManagedPolicyDto `json:"attached_managed_policies,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListManagedPoliciesInPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedPoliciesInPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"ListManagedPoliciesInPermissionSetResponse", string(data)}, " ")
}
