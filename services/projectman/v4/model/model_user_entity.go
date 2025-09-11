package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserEntity 用户信息
type UserEntity struct {

	// 用户Id
	Id *string `json:"id,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o UserEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserEntity struct{}"
	}

	return strings.Join([]string{"UserEntity", string(data)}, " ")
}
