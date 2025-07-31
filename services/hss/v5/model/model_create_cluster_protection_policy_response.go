package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterProtectionPolicyResponse Response Object
type CreateClusterProtectionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateClusterProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterProtectionPolicyResponse", string(data)}, " ")
}
