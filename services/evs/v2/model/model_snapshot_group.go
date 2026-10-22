package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SnapshotGroup 创建快照一致性组。
type SnapshotGroup struct {

	// **参数解释** 云服务器id。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	ServerId *string `json:"server_id,omitempty"`

	// **参数解释** 云硬盘id列表。 **约束限制** 格式为ids=[\"id1\",\"id2\",...,\"idx\"]。 如果设置了server_id，则该参数只能设置指定实例内已挂载的云硬盘，不再支持设置多个跨实例的云硬盘ID。 **取值范围** 不涉及。 **默认取值** 不涉及。
	VolumeIds *[]string `json:"volume_ids,omitempty"`

	// **参数解释** 是否开启快照极速可用。 **约束限制** 不涉及。 **取值范围** - true：开启。 - false：关闭。 **默认取值** false
	InstantAccess *bool `json:"instant_access,omitempty"`

	// **参数解释** 快照一致性组名称。 **约束限制** 最大支持255个字节。 **取值范围** 最小长度：0 最大长度：255 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 快照一致性组描述。 **约束限制** 最大支持255个字节。 **取值范围** 最小长度：0 最大长度：255 **默认取值** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释** 快照一致性组的标签。 标签的value：String类型，最大长度43个字符，可以为空字符串。 由英文字母、数字、下划线、点、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。  标签的value：String类型，最大长度43个字符，可以为空字符串。 由英文字母、数字、下划线、点、中划线、UNICODE字符（\\u4E00-\\u9FFF）组成。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Tags map[string]string `json:"tags,omitempty"`

	// **参数解释** 快照一致性组所属的企业项目ID。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 是否创建增量快照。 **约束限制** 不涉及。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Incremental *bool `json:"incremental,omitempty"`
}

func (o SnapshotGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SnapshotGroup struct{}"
	}

	return strings.Join([]string{"SnapshotGroup", string(data)}, " ")
}
