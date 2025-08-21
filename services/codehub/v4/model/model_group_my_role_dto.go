package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupMyRoleDto struct {

	// 成员id
	Id *int32 `json:"id,omitempty"`

	// 成员角色值
	AccessLevel *int32 `json:"access_level,omitempty"`

	// 角色中文名称
	RoleNamecn *string `json:"role_namecn,omitempty"`

	// 角色名称
	RoleNamen *string `json:"role_namen,omitempty"`

	// 来源代码组id
	SourceId *int32 `json:"source_id,omitempty"`

	// 来源类型
	SourceType *string `json:"source_type,omitempty"`

	// 用户id
	UserId *int32 `json:"user_id,omitempty"`

	// 提示级别
	NotificationLevel *int32 `json:"notification_level,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 创建者id
	CreatedById *int32 `json:"created_by_id,omitempty"`

	// 邀请邮箱
	InviteEmail *string `json:"invite_email,omitempty"`

	// 邀请token
	InviteToken *string `json:"invite_token,omitempty"`

	// 邀请接受时间
	InviteAcceptedAt *string `json:"invite_accepted_at,omitempty"`

	// 请求时间
	RequestedAt *string `json:"requested_at,omitempty"`

	// 过期时间
	ExpiresAt *string `json:"expires_at,omitempty"`

	// 限制
	Limited *bool `json:"limited,omitempty"`

	// 是否为项目管理员
	IsProjectAdmin *int32 `json:"isProjectAdmin,omitempty"`

	// 是否为组织创建者
	IsGroupCreator *int32 `json:"isGroupCreator,omitempty"`

	// 是否仓库创建者
	IsRepoCreator *int32 `json:"isRepoCreator,omitempty"`

	// 展示标记
	RoleShowFlag *int32 `json:"roleShowFlag,omitempty"`
}

func (o GroupMyRoleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupMyRoleDto struct{}"
	}

	return strings.Join([]string{"GroupMyRoleDto", string(data)}, " ")
}
