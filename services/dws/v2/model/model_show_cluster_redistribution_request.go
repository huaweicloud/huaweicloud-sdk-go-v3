package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterRedistributionRequest Request Object
type ShowClusterRedistributionRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页查询，每页大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页偏移量，从0开始，页数减1。 **约束限制**： 不涉及。 **取值范围**： 大于等于0 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 数据库名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**： 表名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： null
	TableName *string `json:"table_name,omitempty"`

	// **参数解释**： 类型，取值来自public.pgxc_redistb表的redistributed字段。多个可用逗号分割。 **约束限制**： 不涉及。 **取值范围**： i：重分布中； y：已完成； n：未开始； **默认取值**： null，即不过滤。
	Type *string `json:"type,omitempty"`
}

func (o ShowClusterRedistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRedistributionRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterRedistributionRequest", string(data)}, " ")
}
