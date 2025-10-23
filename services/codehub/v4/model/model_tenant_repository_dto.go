package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantRepositoryDto 租户下仓库列表
type TenantRepositoryDto struct {

	// **参数解释：** 仓库所有者。 **取值范围：** 字符串长度不少于1，不超过128。
	Owner *string `json:"owner,omitempty"`

	// **参数解释：** 仓库容量,单位:MB,保留2位小数。 **取值范围：** 不涉及。
	Capacity *float64 `json:"capacity,omitempty"`

	// **参数解释：** 仓库状态。 **取值范围：** - 0，正常。 - 3，冻结。 - 4，关闭。 - 5，清理中。 - 7，删除中。
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 内容审核结果。 **取值范围：** true，审核通过。 false，审核未通过。
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释：** 成员数量。 **取值范围：** 不涉及。
	MemberNumber *int32 `json:"member_number,omitempty"`

	// **参数解释：** 仓库Id。 **取值范围：** 不涉及。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 仓库名。 **取值范围：** 不涉及。
	RepositoryName *string `json:"repository_name,omitempty"`

	// **参数解释：** 项目名。 **取值范围：** 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 项目Id。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`
}

func (o TenantRepositoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantRepositoryDto struct{}"
	}

	return strings.Join([]string{"TenantRepositoryDto", string(data)}, " ")
}
