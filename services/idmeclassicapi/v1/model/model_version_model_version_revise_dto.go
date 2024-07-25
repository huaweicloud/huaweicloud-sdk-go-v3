package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionReviseDto struct {

	// **参数解释：**  创建人。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  关系实体名称集合，与workCopyType的值CUSTOM配合使用。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	// **参数解释：**  主对象ID。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  更新者。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  关系的复制类型。  **约束限制：**  不涉及。  **取值范围：**  - BOTH：若存在关系实例引用此数据实例作为源端实例或目标端实例，修订后的数据实例将继承这些关系实例。 - SOURCE：若存在关系实例引用此数据实例作为源端实例，修订后的数据实例将继承这些关系实例。 - TARGET：若存在关系实例引用此数据实例作为目标端实例，修订后的数据实例将继承这些关系实例。 - NONE：修订后的数据实例将不继承任何关系实例。 - CUSTOM：若指定的关系实体集合对应的关系实例引用此数据实例作为源端实例或目标端实例，修订后的数据实例将继承这些关系实例。  **默认取值：**  不涉及。
	WorkCopyType *string `json:"workCopyType,omitempty"`

	// **参数解释：**  是否已检出。  **约束限制：**  不涉及。  **取值范围：**  - true：已检出。 - false：未检出。  **默认取值：**  false。
	WorkingCopy *bool `json:"workingCopy,omitempty"`
}

func (o VersionModelVersionReviseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionReviseDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionReviseDto", string(data)}, " ")
}
