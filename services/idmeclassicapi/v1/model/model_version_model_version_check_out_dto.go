package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionCheckOutDto struct {

	// **参数解释：**  主对象ID，用于定位需要检出的M-V模型实例所属的主对象。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  关系的复制类型，控制检出后的工作副本是否继承原实例的关系实例。不同取值对应不同的关系继承策略，适用于不同的业务场景。  **约束限制：**  不涉及。  **取值范围：**  - BOTH：若存在关系实例引用此数据实例作为源端实例或目标端实例，检出后的数据实例将继承这些关系实例。 - SOURCE：若存在关系实例引用此数据实例作为源端实例，检出后的数据实例将继承这些关系实例。 - TARGET：若存在关系实例引用此数据实例作为目标端实例，检出后的数据实例将继承这些关系实例。 - NONE：检出后的数据实例将不继承任何关系实例。 - CUSTOM：若指定的关系实体集合对应的关系实例引用此数据实例作为源端实例或目标端实例，检出后的数据实例将继承这些关系实例。  **默认取值：**  不涉及。
	WorkCopyType *string `json:"workCopyType,omitempty"`

	// **参数解释：**  关系实体名称集合，与workCopyType的值CUSTOM配合使用。 当需要仅继承部分特定关系实例时，在此指定关系实体的英文名称列表。  **约束限制：**  仅在workCopyType为CUSTOM时生效。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	// **参数解释：**  创建者账号，记录本次生成的工作副本的创建者信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  更新者账号，记录本次检出操作的操作人信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionCheckOutDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionCheckOutDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionCheckOutDto", string(data)}, " ")
}
