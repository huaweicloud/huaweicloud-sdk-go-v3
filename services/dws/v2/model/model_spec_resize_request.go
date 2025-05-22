package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SpecResizeRequest **参数解释**： 规格变更请求信息。 **取值范围**： 不涉及。
type SpecResizeRequest struct {

	// **参数解释**： 目标规格ID。 **取值范围**： 不涉及。
	TargetFlavorId string `json:"target_flavor_id"`

	// **参数解释**： 强制备份。字段已废弃，不再生效。 **取值范围**： 不涉及。
	ForceBackup *bool `json:"force_backup,omitempty"`
}

func (o SpecResizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpecResizeRequest struct{}"
	}

	return strings.Join([]string{"SpecResizeRequest", string(data)}, " ")
}
