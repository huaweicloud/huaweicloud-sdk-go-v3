package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2ProjectTagsRequest Request Object
type ListV2ProjectTagsRequest struct {

	// **参数解释**：端点引擎类型列表。 **约束限制**：不能为空。 **取值范围**：   - aidatalake-workspace：工作空间。   - aidatalake-endpoint-ray：Ray引擎端点。   - aidatalake-endpoint-aura：Aura引擎端点。   - aidatalake-endpoint-spark：Spark引擎端点。   - aidatalake-endpoint-flink：Flink引擎端点。   - aidatalake-compute：计算资源。 **默认取值**：不涉及。
	ResourceType *TagResourceType `json:"resource_type"`

	// **参数解释**：使用预定义标签。 **约束限制**：不涉及。 **取值范围**：   - true：使用。   - false：不使用。 **默认取值**：false。
	UsePredefineTags bool `json:"use_predefine_tags"`

	// **参数解释**：端点引擎类型列表。取值包括：   - AuraJob：AuraJob引擎。   - SparkSql：SparkSql引擎。   - SparkJob：SparkJob引擎。   - FlinkJob：FlinkJob引擎。   - AuraJobV2：AuraJobV2引擎。   - RayCluster：RayCluster引擎。   - RayJob：RayJob引擎。   - RayService：RayService引擎。 **约束限制**：不涉及。
	EngineTypes *[]EngineType `json:"engine_types,omitempty"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：查询记录数。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：1000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListV2ProjectTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2ProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListV2ProjectTagsRequest", string(data)}, " ")
}
