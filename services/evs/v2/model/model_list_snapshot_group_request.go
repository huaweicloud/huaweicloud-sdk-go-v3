package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotGroupRequest Request Object
type ListSnapshotGroupRequest struct {

	// **参数解释** 指定的快照一致性组ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释** 快照一致性组名称。 **约束限制** 最大支持255个字节。 最小长度：0 最大长度：255 **取值范围** 不涉及。 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 快照一致性组的状态。 **约束限制** 不涉及。 **取值范围** - creating：创建中。 - available：创建成功。 - error：创建失败。 - deleting: 删除中。 - deleted：删除成功。 - error_deleting：删除失败。 **默认取值** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释** 快照一致性组的标签的key。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	TagKey *string `json:"tag_key,omitempty"`

	// **参数解释** 快照一致性组的标签的key/value对，例如[{\"key\":\"bbb\",\"value\":\"bbb\"}]。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Tags *string `json:"tags,omitempty"`

	// **参数解释** 快照一致性组的企业项目ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 返回结果个数限制。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 偏移量。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释** 返回结果按该关键字排序。 **约束限制** 不涉及。 **取值范围** - id：快照一致性组的ID。 - status：快照一致性组的状态。 - created_at：快照一致性组的创建时间。 **默认取值** created_at
	SortKey *string `json:"sort_key,omitempty"`

	// **参数解释** 返回结果按照降序或升序排列。 **约束限制** 不涉及。 **取值范围** - desc：降序。 - asc：升序。 **默认取值** desc
	SortDir *string `json:"sort_dir,omitempty"`

	// **参数解释** 云服务器ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	ServerId *string `json:"server_id,omitempty"`
}

func (o ListSnapshotGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotGroupRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotGroupRequest", string(data)}, " ")
}
