package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateImageSnapshotRequestInput **参数解释**：创建镜像快照请求体。
type CreateImageSnapshotRequestInput struct {

	// **参数解释**：快照名称（用户指定，用于标识快照用途）。 **约束限制**：不涉及。 **取值范围**：以小写字母开头，小写字母、数字、中划线（-）组成，小写字母或数字结尾的1到190字符的字符串。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：快照大小，单位GiB。 **约束限制**：不涉及。 **取值范围**：20~400。 **默认取值**：不传则基于镜像大小的3倍自动估算。
	ImageSnapshotSize *int32 `json:"image_snapshot_size,omitempty"`

	// **参数解释**：镜像快照保留天数。超过保留天数自动失效，0表示永久保留。 **约束限制**：不涉及。 **取值范围**：0~65536天。 **默认取值**：0。
	ReservedDays *int32 `json:"reserved_days,omitempty"`
}

func (o CreateImageSnapshotRequestInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateImageSnapshotRequestInput struct{}"
	}

	return strings.Join([]string{"CreateImageSnapshotRequestInput", string(data)}, " ")
}
