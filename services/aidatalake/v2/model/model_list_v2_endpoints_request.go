package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2EndpointsRequest Request Object
type ListV2EndpointsRequest struct {

	// **参数解释**：**参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。  **取值范围**：长度为32-36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：端点引擎类型列表。取值包括：   - AuraJob：AuraJob引擎。   - SparkSql：SparkSql引擎。   - SparkJob：SparkJob引擎。   - FlinkJob：FlinkJob引擎。   - AuraJobV2：AuraJobV2引擎。   - RayCluster：RayCluster引擎。   - RayJob：RayJob引擎。   - RayService：RayService引擎。 **约束限制**：不涉及。
	EngineTypes *[]EngineType `json:"engine_types,omitempty"`

	// **参数解释**：计算资源ID。可通过[查询计算资源列表]获取计算资源ID(ListV2Computes.xml)。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ComputeId *string `json:"compute_id,omitempty"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：端点名称，精确查询。 **约束限制**：不涉及。 **取值范围**：长度为1~63个字符，遵循RFC 1035规范，支持小写字母、数字、中划线，必须字母开头、字母或数字结尾。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：存储资源ID，可通过[查询存储资源列表](ListV2Storages.xml)获取存储资源ID。 **约束限制**：不涉及。 **取值范围**：长度为1~64的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	StorageId *string `json:"storage_id,omitempty"`

	// **参数解释**：标签过滤条件，支持多个标签。格式为 key#value1,value2，用英文逗号分隔多个值。 **约束限制**：最多10个标签。
	Tags *[]string `json:"tags,omitempty"`
}

func (o ListV2EndpointsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2EndpointsRequest struct{}"
	}

	return strings.Join([]string{"ListV2EndpointsRequest", string(data)}, " ")
}
