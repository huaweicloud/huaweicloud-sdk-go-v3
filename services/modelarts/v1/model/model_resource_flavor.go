package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavor 资源规格的数据模型。
type ResourceFlavor struct {

	// **参数解释**：资源的API版本。 **取值范围**：可选值如下： - v1：当前资源版本为v1
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：资源的类型。 **取值范围**：可选值如下： - ResourceFlavor：资源规格
	Kind string `json:"kind"`

	Metadata *ResourceFlavorMetadata `json:"metadata"`

	Spec *ResourceFlavorSpec `json:"spec"`

	Status *ResourceFlavorStatus `json:"status"`
}

func (o ResourceFlavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavor struct{}"
	}

	return strings.Join([]string{"ResourceFlavor", string(data)}, " ")
}
