package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoginPolicyV5Request Request Object
type ShowLoginPolicyV5Request struct {
}

func (o ShowLoginPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoginPolicyV5Request struct{}"
	}

	return strings.Join([]string{"ShowLoginPolicyV5Request", string(data)}, " ")
}
