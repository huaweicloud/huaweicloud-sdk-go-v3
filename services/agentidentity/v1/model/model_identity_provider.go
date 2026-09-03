package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IdentityProvider Identity provider code.
type IdentityProvider struct {
}

func (o IdentityProvider) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdentityProvider struct{}"
	}

	return strings.Join([]string{"IdentityProvider", string(data)}, " ")
}
