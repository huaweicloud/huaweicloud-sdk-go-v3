package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RepositoryUserGroupDto 仓库成员组详情
type RepositoryUserGroupDto struct {

	// **参数解释：** 成员组名称 **约束限制：** 不涉及。
	UserGroupName *string `json:"user_group_name,omitempty"`

	// **参数解释：** 成员组id。 **约束限制：** 不涉及。
	UserGroupId *string `json:"user_group_id,omitempty"`

	// **参数解释：** 项目id。 **约束限制：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 成员组用户数量。 **约束限制：** 不涉及。
	UserCount *string `json:"user_count,omitempty"`

	// **参数解释：** 成员组描述。 **约束限制：** 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o RepositoryUserGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryUserGroupDto struct{}"
	}

	return strings.Join([]string{"RepositoryUserGroupDto", string(data)}, " ")
}
