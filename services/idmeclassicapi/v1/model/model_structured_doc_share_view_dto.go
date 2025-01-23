package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StructuredDocShareViewDto struct {

	// **参数解释**：  唯一标识。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  创建者。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  修改人。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  创建时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  最后更新时间。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  系统版本。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。  **默认取值：**  0。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	StructuredDoc *StructuredDocView `json:"structuredDoc,omitempty"`

	// **参数解释**：  分享用户名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ShareUserName *string `json:"shareUserName,omitempty"`

	// **参数解释**：  被分享用户名。  **约束限制**：  不涉及。  **取值范围**：  - all: 所有人。  **默认取值**：  不涉及。
	SharedUserName *string `json:"sharedUserName,omitempty"`

	// **参数解释**：  被分享用户UserId。  **约束限制**：  不涉及。  **取值范围**：  - all: 所有人。  **默认取值**：  不涉及。
	SharedUserId *string `json:"sharedUserId,omitempty"`

	// **参数解释**：  认证类型。  **约束限制**：  不涉及。  **取值范围**：  - read: 只读 - write: 读写  **默认取值**：  不涉及。
	AuthType *string `json:"authType,omitempty"`

	// **参数解释**：  被分享用户UserId。  **约束限制**：  不涉及。  **取值范围**：  - all: 所有人。  **默认取值**：  不涉及。
	ShareUserId *string `json:"shareUserId,omitempty"`
}

func (o StructuredDocShareViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructuredDocShareViewDto struct{}"
	}

	return strings.Join([]string{"StructuredDocShareViewDto", string(data)}, " ")
}
