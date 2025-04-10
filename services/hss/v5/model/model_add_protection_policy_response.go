package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddProtectionPolicyResponse Response Object
type AddProtectionPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"AddProtectionPolicyResponse", string(data)}, " ")
}
