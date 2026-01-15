package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiOpsRequest Request Object
type ListAiOpsRequest struct {

	// **参数解释**： 指定查询的集群ID。获取方法请参见[获取集群ID](css_03_0101.xml)。 **约束限制**： 不涉及 **取值范围**： 集群ID。 **默认取值**： 不涉及
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 分页参数，列表当前分页的数量限制。默认值为10，即一次查询10个任务信息。 **约束限制**： 不涉及 **取值范围**： 1-1000 **默认取值**： 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，表示从偏移量后面的计数开始查询。 **约束限制**： 不涉及 **取值范围**： 0-1000 **默认取值**： 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 获取当前最新一份报告或历史报告 **约束限制**： 不涉及 **取值范围**： - current   仅获取当前最新一次检测报告 - history   仅获取当前历史检测报告  **默认取值**： 不涉及
	Report *string `json:"report,omitempty"`
}

func (o ListAiOpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiOpsRequest struct{}"
	}

	return strings.Join([]string{"ListAiOpsRequest", string(data)}, " ")
}
