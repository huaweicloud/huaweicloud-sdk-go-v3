package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolResourceFlavorExtendParams **参数解释**：自定义配置参数。 **约束限制**：不涉及。
type PoolResourceFlavorExtendParams struct {

	// **参数解释**：指定资源池节点的容器引擎空间大小。值为0时表示不限制大小。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DockerBaseSize *string `json:"dockerBaseSize,omitempty"`
}

func (o PoolResourceFlavorExtendParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolResourceFlavorExtendParams struct{}"
	}

	return strings.Join([]string{"PoolResourceFlavorExtendParams", string(data)}, " ")
}
