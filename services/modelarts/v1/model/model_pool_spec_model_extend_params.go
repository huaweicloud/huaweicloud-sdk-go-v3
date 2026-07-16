package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecModelExtendParams **参数解释**：资源池自定义配置参数。
type PoolSpecModelExtendParams struct {

	// **参数解释**：资源池创建的节点的容器引擎空间大小。值为0时表示不限制大小。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	DockerBaseSize *string `json:"dockerBaseSize,omitempty"`

	// **参数描述**：磁盘分组名称。 **取值范围**：不涉及。
	VolumeGroup *string `json:"volumeGroup,omitempty"`

	// **参数描述**：模型运行时环境。 **取值范围**：不涉及。
	Runtime *string `json:"runtime,omitempty"`
}

func (o PoolSpecModelExtendParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModelExtendParams struct{}"
	}

	return strings.Join([]string{"PoolSpecModelExtendParams", string(data)}, " ")
}
