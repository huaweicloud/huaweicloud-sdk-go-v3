package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HistoryDataModelHistoryViewDto struct {

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  请求数据。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  修改时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  系统版本。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  操作类型。  **取值范围：**  - CREATE：创建操作。 - UPDATE：更新操作。 - LOGICALDELETE：软删除操作。 - DELETE：删除操作。 - CASCADE：级联操作。  **默认取值：**  不涉及。
	RdmOperationType *string `json:"rdmOperationType,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  删除标志。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	Tenant *TenantHistoryViewDto `json:"tenant,omitempty"`

	// **参数解释：**  类名称。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`
}

func (o HistoryDataModelHistoryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryDataModelHistoryViewDto struct{}"
	}

	return strings.Join([]string{"HistoryDataModelHistoryViewDto", string(data)}, " ")
}
