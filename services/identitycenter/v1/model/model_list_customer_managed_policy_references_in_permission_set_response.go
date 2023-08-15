package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerManagedPolicyReferencesInPermissionSetResponse Response Object
type ListCustomerManagedPolicyReferencesInPermissionSetResponse struct {

	// 指定附加到权限集的客户管理策略的名称和路径.
	CustomerManagedPolicyReferences *[]CustomerManagedPolicyReferenceDto `json:"customer_managed_policy_references,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListCustomerManagedPolicyReferencesInPermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerManagedPolicyReferencesInPermissionSetResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerManagedPolicyReferencesInPermissionSetResponse", string(data)}, " ")
}
