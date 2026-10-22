package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteV2ResourceTagsRequest Request Object
type BatchDeleteV2ResourceTagsRequest struct {

	// **参数解释**：端点引擎类型列表。 **约束限制**：不能为空。 **取值范围**：   - aidatalake-workspace：工作空间。   - aidatalake-endpoint-ray：Ray引擎端点。   - aidatalake-endpoint-aura：Aura引擎端点。   - aidatalake-endpoint-spark：Spark引擎端点。   - aidatalake-endpoint-flink：Flink引擎端点。   - aidatalake-compute：计算资源。 **默认取值**：不涉及。
	ResourceType *TagResourceType `json:"resource_type"`

	// **参数解释**：标签关联的资源ID。 **约束限制**：不能为空。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ResourceId string `json:"resource_id"`

	Body *V2BatchDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteV2ResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteV2ResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteV2ResourceTagsRequest", string(data)}, " ")
}
