package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenericLinkTypeDto struct {

	// **参数解释：**  是否返回源模型数据实例关联的最新版本目标模型数据实例。此参数仅对源/目标模型为M-V模型实体有效。  **约束限制：**  不涉及。  **取值范围：**  - true：返回源模型数据实例关联的最新版本的目标模型数据实例。 - false：返回源模型数据实例关联的所有版本的目标模型数据实例。默认为false。  **默认取值：**  false。
	LatestOnly *bool `json:"latestOnly,omitempty"`

	// **参数解释：**  源模型数据实例的ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	SourceId *string `json:"sourceId,omitempty"`

	// **参数解释：**  目标模型的英文名称。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetType *string `json:"targetType,omitempty"`
}

func (o GenericLinkTypeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericLinkTypeDto struct{}"
	}

	return strings.Join([]string{"GenericLinkTypeDto", string(data)}, " ")
}
