package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistableModelViewDto struct {

	// **参数解释：**  类名，表示数据实例的具体类类型。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者账号。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  数据实例的唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  软删除标识，用于标识实例是否已被逻辑删除。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  扩展类型，表示数据模型的具体扩展类型。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  系统版本号，用于标识实例的版本信息。  **取值范围：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  更新者账号。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *interface{} `json:"lastUpdateTime,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  数据模型的唯一键属性。  **取值范围：**  不涉及。
	UniqueKey *string `json:"uniqueKey,omitempty"`
}

func (o PersistableModelViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistableModelViewDto struct{}"
	}

	return strings.Join([]string{"PersistableModelViewDto", string(data)}, " ")
}
