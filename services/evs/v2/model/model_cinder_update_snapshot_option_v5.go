package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CinderUpdateSnapshotOptionV5 struct {

	// **参数解释** 云硬盘快照名称。最大支持64个字符。 **约束限制** 对云硬盘创建备份时，同时会创建以autobk_snapshot_为名称前缀的快照，云硬盘控制台对此类快照会有操作限制。因此建议不要创建以autobk_snapshot_为名称前缀的快照，避免影响快照的正常使用。 **取值范围** 不涉及。 **默认取值** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 云硬盘快照描述。 **约束限制** 不涉及。 **取值范围** 最大支持85个字符。 **默认取值** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释** 云硬盘快照极速可用标识。 **约束限制** 不涉及。 **取值范围** - true：开启快照极速可用功能。 - false：关闭快照极速可用功能。 **默认取值** 不涉及。
	InstantAccess *bool `json:"instant_access,omitempty"`
}

func (o CinderUpdateSnapshotOptionV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderUpdateSnapshotOptionV5 struct{}"
	}

	return strings.Join([]string{"CinderUpdateSnapshotOptionV5", string(data)}, " ")
}
