package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigIamUserDto IAM用户信息
type ApigIamUserDto struct {

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty"`

	// 是否是空间拥有者
	IsDomainOwner *bool `json:"is_domain_owner,omitempty"`
}

func (o ApigIamUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigIamUserDto struct{}"
	}

	return strings.Join([]string{"ApigIamUserDto", string(data)}, " ")
}
