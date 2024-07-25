package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelVersionCheckoutAndUpdateDtoVersionModel struct {

	// **参数解释：**  创建人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  关系实体名称集合，与workCopyType的值CUSTOM配合使用。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	Data *VersionModel `json:"data"`

	// **参数解释：**  主对象ID。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	MasterId string `json:"masterId"`

	// **参数解释：**  更新者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  关系的复制类型。  **取值范围：**  - BOTH：若存在关系实例引用此数据实例作为源端实例或目标端实例，检出并更新后的数据实例将继承这些关系实例。 - SOURCE：若存在关系实例引用此数据实例作为源端实例，检出并更新后的数据实例将继承这些关系实例。 - TARGET：若存在关系实例引用此数据实例作为目标端实例，检出并更新后的数据实例将继承这些关系实例。 - NONE：检出并更新后的数据实例将不继承任何关系实例。 - CUSTOM：若指定的关系实体集合对应的关系实例引用此数据实例作为源端实例或目标端实例，检出并更新后的数据实例将继承这些关系实例。  **默认取值：**  不涉及。
	WorkCopyType *string `json:"workCopyType,omitempty"`
}

func (o VersionModelVersionCheckoutAndUpdateDtoVersionModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelVersionCheckoutAndUpdateDtoVersionModel struct{}"
	}

	return strings.Join([]string{"VersionModelVersionCheckoutAndUpdateDtoVersionModel", string(data)}, " ")
}
