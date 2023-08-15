package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthScope
type AuthScope struct {
	Domain *AuthScopeDomain `json:"domain,omitempty"`

	Project *AuthScopeProject `json:"project,omitempty"`
}

func (o AuthScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthScope struct{}"
	}

	return strings.Join([]string{"AuthScope", string(data)}, " ")
}
