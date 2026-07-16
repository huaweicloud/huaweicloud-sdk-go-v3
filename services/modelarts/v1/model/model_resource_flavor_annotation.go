package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceFlavorAnnotation 资源规格metadata的注释信息。
type ResourceFlavorAnnotation struct {

	// **参数解释**：资源规格支持的私有镜像的过滤条件。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsFlavorImageFilter *string `json:"os.modelarts.flavor/image.filter,omitempty"`
}

func (o ResourceFlavorAnnotation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceFlavorAnnotation struct{}"
	}

	return strings.Join([]string{"ResourceFlavorAnnotation", string(data)}, " ")
}
