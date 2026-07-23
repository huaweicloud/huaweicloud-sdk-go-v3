package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenericLinkTypeDto struct {

	// **参数解释：**  源模型数据实例的ID，用于查询该实例关联的目标模型数据实例。  **约束限制：**  如不传入，将查询该源模型下所有实例与目标模型的关联关系，返回数据量可能较大。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	SourceId *string `json:"sourceId,omitempty"`

	// **参数解释：**  目标模型的英文名称，用于限定查询源实例与哪个目标模型的关联关系。  **约束限制：**  如不传入，将查询源实例与所有目标模型的关联关系。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetType *string `json:"targetType,omitempty"`

	// **参数解释：**  是否仅返回源实例关联的最新版本目标模型数据实例。此参数仅当源模型或目标模型为M-V模型实体时生效。  **约束限制：**  仅对M-V模型实体有效。  **取值范围：**  - true：仅返回源实例关联的最新版本目标模型数据实例。 - false：返回源实例关联的所有版本目标模型数据实例。  **默认取值：**  false。
	LatestOnly *bool `json:"latestOnly,omitempty"`
}

func (o GenericLinkTypeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericLinkTypeDto struct{}"
	}

	return strings.Join([]string{"GenericLinkTypeDto", string(data)}, " ")
}
