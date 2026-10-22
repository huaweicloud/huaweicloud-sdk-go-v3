package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2ResourceTagsRequest Request Object
type ListV2ResourceTagsRequest struct {

	// **参数解释**：端点引擎类型列表。 **约束限制**：不能为空。 **取值范围**：   - aidatalake-workspace：工作空间。   - aidatalake-endpoint-ray：Ray引擎端点。   - aidatalake-endpoint-aura：Aura引擎端点。   - aidatalake-endpoint-spark：Spark引擎端点。   - aidatalake-endpoint-flink：Flink引擎端点。   - aidatalake-compute：计算资源。 **默认取值**：不涉及。
	ResourceType *TagResourceType `json:"resource_type"`

	// **参数解释**：标签关联的资源ID。 **约束限制**：不能为空。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ResourceId string `json:"resource_id"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：查询记录数。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：1000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListV2ResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2ResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListV2ResourceTagsRequest", string(data)}, " ")
}
