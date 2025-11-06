package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectUserGroupDto struct {

	// **参数解释：** 用户组记录id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户组id。 **取值范围：** 字符串长度不少于1，不超过1000。
	UserGroupId *string `json:"user_group_id,omitempty"`

	// **参数解释：** 用户组名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户组所在项目id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 用户组所在租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 用户组类型。 **取值范围：** 字符串长度不少于1，不超过1000。
	GroupType *string `json:"group_type,omitempty"`

	// **参数解释：** 用户组成员数量。
	MemberCount *int32 `json:"member_count,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 描述信息。 **取值范围：** 字符串长度不少于1，不超过1000。
	Description *string `json:"description,omitempty"`
}

func (o ProjectUserGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectUserGroupDto struct{}"
	}

	return strings.Join([]string{"ProjectUserGroupDto", string(data)}, " ")
}
