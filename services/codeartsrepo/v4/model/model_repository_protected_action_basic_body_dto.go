package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryProtectedActionBasicBodyDto struct {

	// **参数解释：** 是否启用。 **约束限制：** 不涉及。 **取值范围：** - true，开启规则限制 - false，关闭规则限制 **默认取值：** 不涉及。
	Enable *bool `json:"enable,omitempty"`

	// **参数解释：** 用户ID列表。 **约束限制：** 不涉及。 **取值范围：** Integer **默认取值：** 不涉及。
	UserIds *[]int32 `json:"user_ids,omitempty"`

	// **参数解释：** 成员组ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	UserTeamIds *[]int32 `json:"user_team_ids,omitempty"`

	// **参数解释：** 关联角色ID列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	RelatedRoleIds *[]string `json:"related_role_ids,omitempty"`
}

func (o RepositoryProtectedActionBasicBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryProtectedActionBasicBodyDto struct{}"
	}

	return strings.Join([]string{"RepositoryProtectedActionBasicBodyDto", string(data)}, " ")
}
