package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountSnapshotsGroupRequest Request Object
type CountSnapshotsGroupRequest struct {

	// **参数解释** 快照一致性组的名称。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 快照一致性组的状态。 **约束限制** 不涉及。 **取值范围** - creating：创建中。 - available：创建成功。 - error：创建失败。 - deleting: 删除中。 - deleted：删除成功。 - error_deleting：删除失败。 **默认取值** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释** 指定的快照一致性组ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释** 快照一致性组的企业项目ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 快照一致性组的标签的key。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	TagKey *string `json:"tag_key,omitempty"`

	// **参数解释** 快照一致性组的标签的key/value对，例如[{\"key\":\"bbb\",\"value\":\"bbb\"}]。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Tags *string `json:"tags,omitempty"`
}

func (o CountSnapshotsGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountSnapshotsGroupRequest struct{}"
	}

	return strings.Join([]string{"CountSnapshotsGroupRequest", string(data)}, " ")
}
