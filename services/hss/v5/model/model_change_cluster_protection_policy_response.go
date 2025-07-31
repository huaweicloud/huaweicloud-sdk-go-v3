package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeClusterProtectionPolicyResponse Response Object
type ChangeClusterProtectionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeClusterProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeClusterProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ChangeClusterProtectionPolicyResponse", string(data)}, " ")
}
