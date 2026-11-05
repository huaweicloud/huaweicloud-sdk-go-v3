package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeconfigSpecConfigsDropCache **参数解释**：节点缓存清理配置。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type NodeconfigSpecConfigsDropCache struct {

	// **参数解释**：缓存清理开关。 **约束限制**：不涉及。 **取值范围**：enable、disable。 **默认取值**：不涉及。
	Mode *string `json:"mode,omitempty"`
}

func (o NodeconfigSpecConfigsDropCache) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigSpecConfigsDropCache struct{}"
	}

	return strings.Join([]string{"NodeconfigSpecConfigsDropCache", string(data)}, " ")
}
