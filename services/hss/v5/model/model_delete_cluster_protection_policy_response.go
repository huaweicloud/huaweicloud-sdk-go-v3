package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterProtectionPolicyResponse Response Object
type DeleteClusterProtectionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClusterProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterProtectionPolicyResponse", string(data)}, " ")
}
