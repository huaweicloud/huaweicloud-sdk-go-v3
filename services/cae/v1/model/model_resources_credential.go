package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourcesCredential struct {

	// 用户access key，当前access key应该包含访问授权云存储的权限。
	Access *string `json:"access,omitempty"`

	// 用户secret key，当前secret key应该包含访问授权云存储的权限。
	Secret *string `json:"secret,omitempty"`
}

func (o ResourcesCredential) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesCredential struct{}"
	}

	return strings.Join([]string{"ResourcesCredential", string(data)}, " ")
}
