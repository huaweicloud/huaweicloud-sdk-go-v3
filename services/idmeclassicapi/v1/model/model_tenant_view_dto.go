package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TenantViewDto struct {

	// **参数解释：**  租户对象的类名。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  租户编码，用于多租户场景下的租户唯一标识。  **取值范围：**  不涉及。
	Code string `json:"code"`

	// **参数解释：**  租户对象的创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  租户对象的创建者账号。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  租户使用的数据源名称，用于多租户数据隔离。  **取值范围：**  不涉及。
	DataSource string `json:"dataSource"`

	// **参数解释：**  租户描述信息。  **取值范围：**  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：**  租户失效标识，用于标识租户是否已失效。   **取值范围：**  - true：失效。  - false：未失效（默认）。
	DisableFlag *bool `json:"disableFlag,omitempty"`

	// **参数解释：**  租户的唯一标识。   **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  租户KIA密级。   **取值范围：**  不涉及。
	Kiaguid *string `json:"kiaguid,omitempty"`

	// **参数解释：**  租户的最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  租户的更新者账号。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  租户中文名称。  **取值范围：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  租户英文名称。  **取值范围：**  不涉及。
	NameEn *string `json:"nameEn,omitempty"`

	// **参数解释：**  软删除标识。  **取值范围：**  - 0：表示未删除。 - 1：表示已删除。
	RdmDeleteFlag *int32 `json:"rdmDeleteFlag,omitempty"`

	// **参数解释：**  租户扩展类型。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	// **参数解释：**  租户对象的系统版本号。  **取值范围：**  不涉及。
	RdmVersion *int32 `json:"rdmVersion,omitempty"`

	// **参数解释：**  租户安全密级，用于标识租户信息的安全等级。  **取值范围：**  - internal：内部公开。 - secret：秘密。 - confidential：机密。 - top_secret：绝密。
	SecurityLevel *string `json:"securityLevel,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`
}

func (o TenantViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantViewDto struct{}"
	}

	return strings.Join([]string{"TenantViewDto", string(data)}, " ")
}
