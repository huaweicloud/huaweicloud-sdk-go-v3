package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeConfig **参数解释：** 在线服务升级配置。 **约束限制：** 不涉及。
type UpgradeConfig struct {

	// **参数解释：** 升级类型。 **约束限制：** 不涉及。 **取值范围：** - ROLLING：滚动升级，默认值。 **默认取值：** 不涉及。
	Type *string `json:"type,omitempty"`

	RollingUpdate *RollingUpdate `json:"rolling_update,omitempty"`
}

func (o UpgradeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeConfig struct{}"
	}

	return strings.Join([]string{"UpgradeConfig", string(data)}, " ")
}
