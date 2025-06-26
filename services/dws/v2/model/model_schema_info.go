package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemaInfo **参数解释**： 集群模式空间信息。 **取值范围**： 不涉及。
type SchemaInfo struct {

	// **参数解释**： 集群模式名称。 **取值范围**： 不涉及。
	SchemaName string `json:"schema_name"`

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	DatabaseName string `json:"database_name"`

	// **参数解释**： 集群模式使用空间总值。 **取值范围**： 不涉及。
	TotalValue int32 `json:"total_value"`

	// **参数解释**： 集群模式空间阈值。 **取值范围**： 不涉及。
	PermSpace int32 `json:"perm_space"`

	// **参数解释**： 倾斜率。 **取值范围**： 不涉及。
	SkewPercent float64 `json:"skew_percent"`

	// **参数解释**： 最小值。 **取值范围**： 不涉及。
	MinValue int32 `json:"min_value"`

	// **参数解释**： 最大值。 **取值范围**： 不涉及。
	MaxValue int32 `json:"max_value"`

	// **参数解释**： 最小dn节点。 **取值范围**： 不涉及。
	MinDn string `json:"min_dn"`

	// **参数解释**： 最大cn节点。 **取值范围**： 不涉及。
	MaxDn string `json:"max_dn"`

	// **参数解释**： dn节点数量。 **取值范围**： 不涉及。
	DnNum int32 `json:"dn_num"`
}

func (o SchemaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaInfo struct{}"
	}

	return strings.Join([]string{"SchemaInfo", string(data)}, " ")
}
