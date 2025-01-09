package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsCredential obs访问凭据。
type ObsCredential struct {

	// access key。
	Access *string `json:"access,omitempty"`

	// secret key。
	Secret *string `json:"secret,omitempty"`

	// 安全校验token。
	SecurityToken *string `json:"security_token,omitempty"`
}

func (o ObsCredential) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsCredential struct{}"
	}

	return strings.Join([]string{"ObsCredential", string(data)}, " ")
}
