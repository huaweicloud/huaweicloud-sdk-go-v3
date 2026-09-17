package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeStrategy **参数解释：** 升级配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeStrategy struct {

	// **参数解释：** 升级策略类型 **约束限制：** 当前仅支持inPlaceRollingUpdate类型 **取值范围：** - \"inPlaceRollingUpdate\"：原地升级类型  **默认取值：** 不涉及
	Type string `json:"type"`

	InPlaceRollingUpdate *InPlaceRollingUpdate `json:"inPlaceRollingUpdate,omitempty"`
}

func (o UpgradeStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeStrategy struct{}"
	}

	return strings.Join([]string{"UpgradeStrategy", string(data)}, " ")
}
