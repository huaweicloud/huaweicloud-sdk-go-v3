package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoginPolicyV5Response Response Object
type UpdateLoginPolicyV5Response struct {
	LoginPolicy    *LoginPolicy `json:"login_policy,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateLoginPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoginPolicyV5Response struct{}"
	}

	return strings.Join([]string{"UpdateLoginPolicyV5Response", string(data)}, " ")
}
