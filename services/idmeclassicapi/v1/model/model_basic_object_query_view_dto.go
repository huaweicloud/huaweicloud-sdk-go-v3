package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicObjectQueryViewDto struct {

	// **参数解释：**  唯一标识，系统生成的根节点主键ID。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  创建者账号，标识创建该根节点的用户。  **取值范围：**  不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释：**  创建时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	CreateTime *string `json:"createTime,omitempty"`

	// **参数解释：**  更新者账号，标识最后更新该根节点的用户。  **取值范围：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`

	// **参数解释：**  最后的修改时间。使用UTC+0时间格式，格式为yyyy-MM-ddTHH:mm:ss.SSSZ。  **取值范围：**  不涉及。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	// **参数解释：**  扩展类型，标识对象的扩展类别。  **取值范围：**  不涉及。
	RdmExtensionType *string `json:"rdmExtensionType,omitempty"`

	Tenant *TenantViewDto `json:"tenant,omitempty"`

	// **参数解释：**  类名，标识对象的Java类名称。  **取值范围：**  不涉及。
	ClassName *string `json:"className,omitempty"`
}

func (o BasicObjectQueryViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicObjectQueryViewDto struct{}"
	}

	return strings.Join([]string{"BasicObjectQueryViewDto", string(data)}, " ")
}
