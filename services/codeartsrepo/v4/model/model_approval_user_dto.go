package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalUserDto struct {

	// 用户id
	Id *int32 `json:"id,omitempty"`

	// 用户名称
	Username *string `json:"username,omitempty"`

	// 用户名称
	Name *string `json:"name,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 中文名称
	NameCn *string `json:"name_cn,omitempty"`

	// 电子邮箱
	Email *string `json:"email,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 头像url
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// 审核意见
	ApproverComment *string `json:"approver_comment,omitempty"`
}

func (o ApprovalUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalUserDto struct{}"
	}

	return strings.Join([]string{"ApprovalUserDto", string(data)}, " ")
}
