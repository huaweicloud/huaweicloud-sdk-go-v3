package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolDriverStatus 资源池驱动状态数据模型。
type PoolDriverStatus struct {

	// **参数解释**：资源池当前驱动版本。 **取值范围**：不涉及。
	Version string `json:"version"`

	// **参数解释**：资源池当前驱动状态。 **取值范围**：可选值如下： - Creating：驱动安装中。 - Upgrading：驱动升级中。 - Running：驱动运行中。 - Abnormal：驱动异常。
	State string `json:"state"`
}

func (o PoolDriverStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolDriverStatus struct{}"
	}

	return strings.Join([]string{"PoolDriverStatus", string(data)}, " ")
}
