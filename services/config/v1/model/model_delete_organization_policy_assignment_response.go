package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationPolicyAssignmentResponse Response Object
type DeleteOrganizationPolicyAssignmentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOrganizationPolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationPolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationPolicyAssignmentResponse", string(data)}, " ")
}
