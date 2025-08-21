package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepositoryMemberDto 添加仓库成员信息
type CreateRepositoryMemberDto struct {

	// **参数解释：** 用户iamId **约束限制：** 不涉及。
	UserIamId *string `json:"user_iam_id,omitempty"`

	// **参数解释：** 用户名称。 **约束限制：** 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 租户名称。 **约束限制：** - 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 租户id。 **约束限制：** - 不涉及。
	TenantId *string `json:"tenant_id,omitempty"`

	// **参数解释：** 角色id。 **约束限制：** - 不涉及。
	RepositoryRoleId *string `json:"repository_role_Id,omitempty"`
}

func (o CreateRepositoryMemberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepositoryMemberDto struct{}"
	}

	return strings.Join([]string{"CreateRepositoryMemberDto", string(data)}, " ")
}
