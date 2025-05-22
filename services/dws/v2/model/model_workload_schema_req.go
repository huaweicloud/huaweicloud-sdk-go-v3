package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadSchemaReq **参数解释**： 模式空间信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadSchemaReq struct {

	// **参数解释**： 模式空间名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SchemaName string `json:"schema_name"`

	// **参数解释**： 模式空间阈值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PermSpace string `json:"perm_space"`
}

func (o WorkloadSchemaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadSchemaReq struct{}"
	}

	return strings.Join([]string{"WorkloadSchemaReq", string(data)}, " ")
}
