package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorLabel 资源规格的标签信息。
type ResourceFlavorLabel struct {

	// **参数解释**：资源规格支持作业类型，以“.”分割。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsScope string `json:"os.modelarts/scope"`
}

func (o ResourceFlavorLabel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorLabel struct{}"
	}

	return strings.Join([]string{"ResourceFlavorLabel", string(data)}, " ")
}
