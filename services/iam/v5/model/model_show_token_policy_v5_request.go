package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTokenPolicyV5Request Request Object
type ShowTokenPolicyV5Request struct {
}

func (o ShowTokenPolicyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTokenPolicyV5Request struct{}"
	}

	return strings.Join([]string{"ShowTokenPolicyV5Request", string(data)}, " ")
}
