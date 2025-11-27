package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NamedAuthInfo struct {

	// 用户名
	Name *string `json:"name,omitempty"`

	User *AuthInfo `json:"user,omitempty"`
}

func (o NamedAuthInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamedAuthInfo struct{}"
	}

	return strings.Join([]string{"NamedAuthInfo", string(data)}, " ")
}
