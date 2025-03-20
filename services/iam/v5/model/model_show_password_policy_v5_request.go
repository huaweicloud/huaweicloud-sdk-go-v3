package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPasswordPolicyV5Request Request Object
type ShowPasswordPolicyV5Request struct {
}

func (o ShowPasswordPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPasswordPolicyV5Request struct{}"
	}

	return strings.Join([]string{"ShowPasswordPolicyV5Request", string(data)}, " ")
}
