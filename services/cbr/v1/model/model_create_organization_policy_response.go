package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationPolicyResponse Response Object
type CreateOrganizationPolicyResponse struct {
	Policy         *OrganizationPolicy `json:"policy,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o CreateOrganizationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateOrganizationPolicyResponse", string(data)}, " ")
}
