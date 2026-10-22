package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListV2TenantImageVersionsRequest Request Object
type ListV2TenantImageVersionsRequest struct {

	// **参数解释**：镜像ID。可通过[查询自定义镜像列表]获取镜像ID(ListV2TenantImages.xml)。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	ImageId string `json:"image_id"`

	// **参数解释**：分页大小。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：100。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：偏移量，表示从此偏移量开始查询。 **约束限制**：不涉及。 **取值范围**：0~1000000000。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：镜像版本名称（TAG）。 **约束限制**：不涉及。 **取值范围**：长度为2~32的英文字符、数字、中划线、点和下划线的组合。 **默认取值**：不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**：镜像版本类型。 **约束限制**：不涉及。 **取值范围**：   - STABLE：稳定版本。   - RECOMMEND：推荐版本。   - BETA：测试版本。 **默认取值**：不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**：镜像架构，多个架构用逗号分隔。 **约束限制**：不涉及。 **取值范围**：   - x86_64：x86架构。   - aarch64：ARM架构。 **默认取值**：不涉及
	Architectures *string `json:"architectures,omitempty"`

	// **参数解释**：工作空间ID，用于查询过滤工作空间支持的镜像版本。可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为1~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：端点ID，用于查询过滤端点支持的镜像版本。可通过[查询端点列表](ListV2Endpoints.xml)获取端点ID。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：指定排序字段。 **约束限制**：不涉及。 **取值范围**：   - CREATE_TIME：创建时间，使用UTC时间，ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSXXX，例如：2025-01-10T03:37:12.000+00:00。   - UPDATE_TIME：更新时间，使用UTC时间，ISO 8601格式，yyyy-MM-dd'T'HH:mm:ss.SSSXXX，例如：2025-01-10T03:37:12.000+00:00。 **默认取值**：CREATE_TIME。
	SortBy *string `json:"sort_by,omitempty"`

	// **参数解释**：排序方式。 **约束限制**：不涉及。 **取值范围**：   - ASC：递增排序。   - DESC：递减排序。 **默认取值**：DESC。
	OrderBy *string `json:"order_by,omitempty"`

	// **参数解释**：租户ID，用于标识租户。可通过IAM服务或控制台获取租户ID。获取方式请参见[获取账号、IAM用户、项目、用户组、委托的名称和ID](iam_01_0023.xml)。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TenantId *string `json:"tenantId,omitempty"`
}

func (o ListV2TenantImageVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListV2TenantImageVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListV2TenantImageVersionsRequest", string(data)}, " ")
}
