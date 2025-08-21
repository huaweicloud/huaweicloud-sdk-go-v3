package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupMyRoleDtoV4 struct {

	// **参数解释：** 记录id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 角色。
	AccessLevel *int32 `json:"access_level,omitempty"`

	// **参数解释：** 角色中文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleNamecn *string `json:"role_namecn,omitempty"`

	// **参数解释：** 角色英文名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RoleNamen *string `json:"role_namen,omitempty"`

	// **参数解释：** 资源id。
	SourceId *int32 `json:"source_id,omitempty"`

	// **参数解释：** 资源类型。 **取值范围：Group Project** 字符串长度不少于1，不超过1000。
	SourceType *string `json:"source_type,omitempty"`

	// **参数解释：** 用户id。
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释：** 通知。
	NotificationLevel *int32 `json:"notification_level,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 是否是项目创建者。
	IsProjectAdmin *int32 `json:"is_project_admin,omitempty"`

	// **参数解释：** 是否是代码组创建者。
	IsGroupCreator *int32 `json:"is_group_creator,omitempty"`

	// **参数解释：** 是否是仓库创建者。
	IsRepoCreator *int32 `json:"is_repo_creator,omitempty"`

	// **参数解释：** 角色展示。
	RoleShowFlag *int32 `json:"role_show_flag,omitempty"`
}

func (o GroupMyRoleDtoV4) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupMyRoleDtoV4 struct{}"
	}

	return strings.Join([]string{"GroupMyRoleDtoV4", string(data)}, " ")
}
