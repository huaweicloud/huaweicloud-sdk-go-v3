package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthObjectAuth 鉴权信息
type AuthObjectAuth struct {
	AccessKey *AuthObjectAuthAccessKey `json:"accessKey,omitempty"`
}

func (o AuthObjectAuth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthObjectAuth struct{}"
	}

	return strings.Join([]string{"AuthObjectAuth", string(data)}, " ")
}
