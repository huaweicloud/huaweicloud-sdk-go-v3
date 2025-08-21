package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForkRepositoryDto struct {

	// **参数解释：** 仓库ID。 **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库名称。 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 是否归档。 **约束限制：** 不涉及。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 产品ID。 **约束限制：** 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释：** 产品名称。 **约束限制：** 不涉及。
	ProductName *string `json:"product_name,omitempty"`

	// **参数解释：** 带命名空间的仓库路径。 **约束限制：** view=least时返回
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`

	// **参数解释：** 命名空间。 **约束限制：** view=basic时返回
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释：** 仓库路径。 **约束限制：** view=basic时返回
	Path *string `json:"path,omitempty"`

	// **参数解释：** 开发模式。 **约束限制：** view=basic时返回
	DevelopMode *string `json:"develop_mode,omitempty"`

	// **参数解释：** 可见性。 **约束限制：** view=basic时返回
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释：** 安全级别。 **约束限制：**  view=basic时返回
	Security *string `json:"security,omitempty"`

	// **参数解释：** 关注数。 **约束限制：** view=basic时返回
	StarCount *int32 `json:"star_count,omitempty"`

	// **参数解释：** Fork数。 **约束限制：** view=basic时返回
	ForksCount *int32 `json:"forks_count,omitempty"`

	// **参数解释：** 开启的合并请求数。 **约束限制：** view=basic时返回
	OpenMergeRequestsCount *int32 `json:"open_merge_requests_count,omitempty"`

	// **参数解释：** 是否已关注。 **约束限制：** view=basic时返回
	Starred *bool `json:"starred,omitempty"`

	// **参数解释：** 带命名空间的仓库名称。 **约束限制：** view=basic时返回
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// **参数解释：** 最后活动时间。 **约束限制：** view=basic时返回
	LastActivityAt *string `json:"last_activity_at,omitempty"`

	// **参数解释：** 创建时间。 **约束限制：** view=basic时返回
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o ForkRepositoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForkRepositoryDto struct{}"
	}

	return strings.Join([]string{"ForkRepositoryDto", string(data)}, " ")
}
