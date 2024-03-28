package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationPolicyResponse Response Object
type UpdateOrganizationPolicyResponse struct {
	Policy         *OrganizationPolicy `json:"policy,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateOrganizationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationPolicyResponse", string(data)}, " ")
}
