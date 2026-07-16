package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolModel 资源池的详细信息。
type PoolModel struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v2：当前资源版本为v2。
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - Pool：资源池。
	Kind string `json:"kind"`

	Metadata *PoolMetadata `json:"metadata"`

	Spec *PoolSpecModel `json:"spec"`

	Status *PoolStatus `json:"status"`
}

func (o PoolModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolModel struct{}"
	}

	return strings.Join([]string{"PoolModel", string(data)}, " ")
}
