package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2TenantImagesRequest Request Object
type ListV2TenantImagesRequest struct {

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：分页大小。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：指定排序字段。 **约束限制**：不涉及。 **取值范围**：   - CREATE_TIME：创建时间，使用UTC时间，ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSXXX，例如：2025-01-10T03:37:12.000+00:00。   - UPDATE_TIME：更新时间，使用UTC时间，ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSXXX，例如：2025-01-10T03:37:12.000+00:00。 **默认取值**：CREATE_TIME。
	SortBy *string `json:"sort_by,omitempty"`

	// **参数解释**：排序方式。 **约束限制**：不涉及。 **取值范围**：   - ASC：递增排序。   - DESC：递减排序。 **默认取值**：DESC。
	OrderBy *string `json:"order_by,omitempty"`

	// **参数解释**：镜像id检索。可通过[查询自定义镜像列表](ListV2TenantImages.xml)获取镜像ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：通过自定义镜像分类检索。 **约束限制**：不涉及。 **取值范围**：   - RAY_BASIC：基于Ray构建的自定义镜像。   - SPARK_JOB：基于SparkJob构建的自定义镜像。   - SPARK_SQL：基于SparkSql构建的自定义镜像。   - FLINK_JOB：基于FlinkJob构建的自定义镜像。   - AURA_JOB：基于AuraJobV2构建的自定义镜像。 **默认取值**：不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**：通过镜像组织搜索。 **约束限制**：不涉及。 **取值范围**：由英文字母、数字、中划线和下划线组成，长度为0~64个字符。 **默认取值**：不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**：通过镜像名搜索。 **约束限制**：不涉及。 **取值范围**：由中文字符、英文字母、数字、中划线和下划线组成，长度为0~128个字符。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：可见性检索的参数。 **约束限制**：不涉及。 **取值范围**：   - PRIVATE：私有，用户自己创建的。   - PUBLIC：公共，查询所有公共的，包括其他用户创建的。   - ALL：所有的。 **默认取值**：默认为空，不填表示不限制。当前用户下包括PRIVATE和PUBLIC的，不包括其他用户创建的。
	Visibility *string `json:"visibility,omitempty"`
}

func (o ListV2TenantImagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2TenantImagesRequest struct{}"
	}

	return strings.Join([]string{"ListV2TenantImagesRequest", string(data)}, " ")
}
