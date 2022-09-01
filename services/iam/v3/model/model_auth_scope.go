package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AuthScope struct {
	Domain *AuthScopeDomain `json:"domain,omitempty" xml:"domain"`

	Project *AuthScopeProject `json:"project,omitempty" xml:"project"`
}

func (o AuthScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthScope struct{}"
	}

	return strings.Join([]string{"AuthScope", string(data)}, " ")
}
