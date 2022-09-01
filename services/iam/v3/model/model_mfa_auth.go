package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type MfaAuth struct {
	Identity *MfaIdentity `json:"identity" xml:"identity"`

	Scope *AuthScope `json:"scope" xml:"scope"`
}

func (o MfaAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MfaAuth struct{}"
	}

	return strings.Join([]string{"MfaAuth", string(data)}, " ")
}
