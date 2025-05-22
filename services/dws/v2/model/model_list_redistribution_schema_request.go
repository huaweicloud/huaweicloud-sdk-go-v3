package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRedistributionSchemaRequest Request Object
type ListRedistributionSchemaRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： ^[a-zA-Z0-9\\u4e00-\\u9fa5_.+= :@!#-]{0,255}$ **默认取值**： null
	DbName string `json:"db_name"`

	// **参数解释**： 分页条数。 **约束限制**： 不涉及。 **取值范围**： 有效值：大于等于1。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量。 **约束限制**： 不涉及。 **取值范围**： 有效值：大于等于0。 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： schema名称。 **约束限制**： 不涉及。 **取值范围**： ^[a-zA-Z0-9\\u4e00-\\u9fa5_.+= ,:@!#-]{0,2048}$ **默认取值**： null
	SchemaName *string `json:"schema_name,omitempty"`
}

func (o ListRedistributionSchemaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedistributionSchemaRequest struct{}"
	}

	return strings.Join([]string{"ListRedistributionSchemaRequest", string(data)}, " ")
}
