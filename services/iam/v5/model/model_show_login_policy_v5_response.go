package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoginPolicyV5Response Response Object
type ShowLoginPolicyV5Response struct {
	LoginPolicy    *LoginPolicy `json:"login_policy,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowLoginPolicyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoginPolicyV5Response struct{}"
	}

	return strings.Join([]string{"ShowLoginPolicyV5Response", string(data)}, " ")
}
