package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModuleOwner struct {

	// 用户32位字符串id
	UserId *string `json:"user_id,omitempty"`

	// 用户数字id
	UserNumId *int32 `json:"user_num_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`
}

func (o ModuleOwner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleOwner struct{}"
	}

	return strings.Join([]string{"ModuleOwner", string(data)}, " ")
}
