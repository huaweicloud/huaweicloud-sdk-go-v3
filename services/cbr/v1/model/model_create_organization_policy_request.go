package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationPolicyRequest Request Object
type CreateOrganizationPolicyRequest struct {
	Body *OrganizationPolicyCreateReq `json:"body,omitempty"`
}

func (o CreateOrganizationPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateOrganizationPolicyRequest", string(data)}, " ")
}
