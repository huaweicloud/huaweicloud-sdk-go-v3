package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionModelMasterViewDto struct {

	// **参数解释：**  主对象的类名。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  主对象的创建时间，使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  主对象的创建者账号。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  主对象的唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  主对象的最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  主对象的更新者账号。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  主对象的扩展类型。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  主对象的系统版本号。  **取值范围：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`
}

func (o VersionModelMasterViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionModelMasterViewDto struct{}"
	}

	return strings.Join([]string{"VersionModelMasterViewDto", string(data)}, " ")
}
