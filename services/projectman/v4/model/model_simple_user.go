package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleUser struct {

	// 用户数字id
	UserNumId *int32 `json:"user_num_id,omitempty" xml:"user_num_id"`

	// 用户uuid
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 账号名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`
}

func (o SimpleUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleUser struct{}"
	}

	return strings.Join([]string{"SimpleUser", string(data)}, " ")
}
