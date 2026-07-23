package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenericLinkTypeModifierDto struct {

	// **参数解释：**  源模型数据实例的ID，用于指定删除哪个源实例与目标模型的关联关系。  **约束限制：**  如不传入，将删除该源模型下所有实例与目标模型的关联关系，请谨慎使用。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	SourceId *string `json:"sourceId,omitempty"`

	// **参数解释：**  目标模型的英文名称，用于指定删除源实例与哪个目标模型的关联关系。  **约束限制：**  如不传入，将删除源实例与所有目标模型的关联关系，请谨慎使用。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TargetType *string `json:"targetType,omitempty"`

	// **参数解释：**  是否仅删除源实例关联的最新版本目标模型数据实例的关系。此参数仅当源模型或目标模型为M-V（Master-View）模型实体时生效。  **约束限制：**  仅对M-V模型实体有效。  **取值范围：**  - true：仅删除源实例关联的最新版本目标模型数据实例的关系。 - false：删除源实例关联的所有版本目标模型数据实例的关系。  **默认取值：**  false。
	LatestOnly *bool `json:"latestOnly,omitempty"`

	// **参数解释：**  更新者账号，用于记录删除操作执行者信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o GenericLinkTypeModifierDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericLinkTypeModifierDto struct{}"
	}

	return strings.Join([]string{"GenericLinkTypeModifierDto", string(data)}, " ")
}
