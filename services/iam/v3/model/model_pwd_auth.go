package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PwdAuth
type PwdAuth struct {
	Identity *PwdIdentity `json:"identity"`

	Scope *AuthScope `json:"scope"`
}

func (o PwdAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PwdAuth struct{}"
	}

	return strings.Join([]string{"PwdAuth", string(data)}, " ")
}
