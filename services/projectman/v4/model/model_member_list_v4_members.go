package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MemberListV4Members struct {

	// 租户id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 用户id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 用户名
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 创建人numId
	UserNumId *int32 `json:"user_num_id,omitempty" xml:"user_num_id"`

	// 成员角色, -1 项目创建者, 3 项目经理, 4 开发人员, 5 测试经理, 6 测试人员, 7 参与者, 8 浏览者, 9 运维经理
	RoleId *int32 `json:"role_id,omitempty" xml:"role_id"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	// 用户角色
	RoleName *string `json:"role_name,omitempty" xml:"role_name"`

	// 用户类型, User iam用户, Federation 联邦账号,
	UserType *string `json:"user_type,omitempty" xml:"user_type"`

	// 是否是禁用账号，1 禁用账号， 0非禁用账号
	Forbidden *int32 `json:"forbidden,omitempty" xml:"forbidden"`
}

func (o MemberListV4Members) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberListV4Members struct{}"
	}

	return strings.Join([]string{"MemberListV4Members", string(data)}, " ")
}
