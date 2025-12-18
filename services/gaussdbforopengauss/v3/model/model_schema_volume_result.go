package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SchemaVolumeResult struct {

	// **参数解释**: schema的大小。 **取值范围**: 不涉及。
	SchemaSize *string `json:"schema_size,omitempty"`

	// **参数解释**: schema拥有的表数量。 **取值范围**: 不涉及。
	TableCount *string `json:"table_count,omitempty"`

	// **参数解释**: schema所属用户名称。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: schema名称空间的名称。 **取值范围**: 不涉及。
	NspName *string `json:"nsp_name,omitempty"`
}

func (o SchemaVolumeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaVolumeResult struct{}"
	}

	return strings.Join([]string{"SchemaVolumeResult", string(data)}, " ")
}
