package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserVo struct {

	// 用户的租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 用户的租户名称
	DomainName *string `json:"domain_name,omitempty"`

	// 用户的昵称
	NickName *string `json:"nick_name,omitempty"`

	// 用户Iam id
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 用户索引id
	UserNumId *int32 `json:"user_num_id,omitempty"`
}

func (o UserVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserVo struct{}"
	}

	return strings.Join([]string{"UserVo", string(data)}, " ")
}
