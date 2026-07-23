package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionObjectBranchQueryViewDto struct {

	// **参数解释：**  类名。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  创建者账号。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  唯一标识。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  最后更新时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  更新者账号。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  扩展类型。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantQueryViewDto `json:"tenant,omitempty"`

	// **参数解释：**  版本号。  **取值范围：**  不涉及。
	Version *string `json:"version,omitempty"`
}

func (o VersionObjectBranchQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionObjectBranchQueryViewDto struct{}"
	}

	return strings.Join([]string{"VersionObjectBranchQueryViewDto", string(data)}, " ")
}
