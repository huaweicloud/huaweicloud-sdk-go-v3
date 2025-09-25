package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateProtectionPolicyResponse Response Object
type AssociateProtectionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateProtectionPolicyResponse", string(data)}, " ")
}
