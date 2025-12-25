package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDatabaseAuthorityRequest Request Object
type ShowDatabaseAuthorityRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 对象类型。 **取值范围**： - DATABASE：数据库。 - SCHEMA：模式。 - TABLE：表。 - VIEW：视图。 - COLUMN：列。 - FUNCTION：函数。 - SEQUENCE：序列。 - NODEGROUP：节点组。
	Type string `json:"type"`

	// **参数解释**： 对象名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name []string `json:"name"`

	// **参数解释**： 数据库名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Database string `json:"database"`

	// **参数解释**： 模式名，对象类型为TABLE、VIEW、COLUMN、FUNCTION、SEQUENCE时必选。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Schema *string `json:"schema,omitempty"`

	// **参数解释**： 表名，对象类型为COLUMN时必选。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Table *string `json:"table,omitempty"`
}

func (o ShowDatabaseAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDatabaseAuthorityRequest struct{}"
	}

	return strings.Join([]string{"ShowDatabaseAuthorityRequest", string(data)}, " ")
}
