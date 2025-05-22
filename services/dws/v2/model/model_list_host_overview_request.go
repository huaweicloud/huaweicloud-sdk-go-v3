package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostOverviewRequest Request Object
type ListHostOverviewRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 分页单页大小。 **约束限制**： 不涉及。 **取值范围**： 大于0。 **默认取值**： 不限制。
	Limit int32 `json:"limit"`

	// **参数解释**： 分页查询，偏移量。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0
	Offset int32 `json:"offset"`
}

func (o ListHostOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostOverviewRequest struct{}"
	}

	return strings.Join([]string{"ListHostOverviewRequest", string(data)}, " ")
}
