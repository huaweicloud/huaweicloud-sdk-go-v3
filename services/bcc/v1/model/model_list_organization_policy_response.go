package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationPolicyResponse Response Object
type ListOrganizationPolicyResponse struct {

	// count
	Count *int32 `json:"count,omitempty"`

	// bccOrganizationPolicies
	BccOrganizationPolicies *[]BccOrganizationPolicy `json:"bcc_organization_policies,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListOrganizationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationPolicyResponse", string(data)}, " ")
}
