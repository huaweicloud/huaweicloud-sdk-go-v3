package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectionPolicyResponse Response Object
type DeleteProtectionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectionPolicyResponse", string(data)}, " ")
}
