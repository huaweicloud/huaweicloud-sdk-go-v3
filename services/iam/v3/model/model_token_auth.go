package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TokenAuth
type TokenAuth struct {
	Identity *TokenAuthIdentity `json:"identity"`
}

func (o TokenAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokenAuth struct{}"
	}

	return strings.Join([]string{"TokenAuth", string(data)}, " ")
}
