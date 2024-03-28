package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationPolicyResponse Response Object
type DeleteOrganizationPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOrganizationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationPolicyResponse", string(data)}, " ")
}
