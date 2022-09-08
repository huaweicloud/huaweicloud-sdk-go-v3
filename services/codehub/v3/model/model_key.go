package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Key struct {

	// 是否允许推送
	CanPush *bool `json:"can_push,omitempty"`

	// 部署密钥新建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 部署密钥
	Key *string `json:"key,omitempty"`

	// 部署密钥id
	KeyId *string `json:"key_id,omitempty"`

	// 部署密钥名称
	KeyTitle *string `json:"key_title,omitempty"`
}

func (o Key) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Key struct{}"
	}

	return strings.Join([]string{"Key", string(data)}, " ")
}
