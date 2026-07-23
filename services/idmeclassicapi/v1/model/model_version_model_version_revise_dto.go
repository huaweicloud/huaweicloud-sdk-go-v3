package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionReviseDto struct {

	// **参数解释：**  主对象ID，用于定位需要修订的M-V模型实例所属的主对象。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  关系的复制类型，控制修订后的新版本是否继承原实例的关系实例。不同取值对应不同的关系继承策略，适用于不同的业务场景。  **约束限制：**  不涉及。  **取值范围：**  - BOTH：若存在关系实例引用此数据实例作为源端实例或目标端实例，修订后的数据实例将继承这些关系实例。 - SOURCE：若存在关系实例引用此数据实例作为源端实例，修订后的数据实例将继承这些关系实例。 - TARGET：若存在关系实例引用此数据实例作为目标端实例，修订后的数据实例将继承这些关系实例。 - NONE：修订后的数据实例将不继承任何关系实例。 - CUSTOM：若指定的关系实体集合对应的关系实例引用此数据实例作为源端实例或目标端实例，修订后的数据实例将继承这些关系实例。  **默认取值：**  不涉及。
	WorkCopyType *string `json:"workCopyType,omitempty"`

	// **参数解释：**  关系实体名称集合，与workCopyType的值CUSTOM配合使用。 指定需要继承的关系实体名称列表，仅继承这些关系实体对应的关系实例。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	// **参数解释：**  修订后是否自动检出。若设置为true，修订后的新版本将自动处于已检出状态（生成工作副本），可直接进行编辑；若设置为false（默认），修订后的新版本处于已检入状态。  **约束限制：**  不涉及。  **取值范围：**  - true：修订后自动检出。 - false：修订后不自动检出。  **默认取值：**  false。
	WorkingCopy *bool `json:"workingCopy,omitempty"`

	// **参数解释：**  创建者账号，记录本次生成的新版本的创建者信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  更新者账号，记录本次修订操作的操作人信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o VersionModelVersionReviseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionReviseDto struct{}"
	}

	return strings.Join([]string{"VersionModelVersionReviseDto", string(data)}, " ")
}
