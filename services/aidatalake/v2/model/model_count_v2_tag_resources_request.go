package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountV2TagResourcesRequest Request Object
type CountV2TagResourcesRequest struct {

	// **参数解释**：端点引擎类型列表。 **约束限制**：不能为空。 **取值范围**：   - aidatalake-workspace：工作空间。   - aidatalake-endpoint-ray：Ray引擎端点。   - aidatalake-endpoint-aura：Aura引擎端点。   - aidatalake-endpoint-spark：Spark引擎端点。   - aidatalake-endpoint-flink：Flink引擎端点。   - aidatalake-compute：计算资源。 **默认取值**：不涉及。
	ResourceType *TagResourceType `json:"resource_type"`

	// **参数解释**：端点引擎类型列表。取值包括：   - AuraJob：AuraJob引擎。   - SparkSql：SparkSql引擎。   - SparkJob：SparkJob引擎。   - FlinkJob：FlinkJob引擎。   - AuraJobV2：AuraJobV2引擎。   - RayCluster：RayCluster引擎。   - RayJob：RayJob引擎。   - RayService：RayService引擎。 **约束限制**：不涉及。
	EngineTypes *[]EngineType `json:"engine_types,omitempty"`

	Body *V2ListTagResourcesRequestBody `json:"body,omitempty"`
}

func (o CountV2TagResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountV2TagResourcesRequest struct{}"
	}

	return strings.Join([]string{"CountV2TagResourcesRequest", string(data)}, " ")
}
