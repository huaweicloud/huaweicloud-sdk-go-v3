package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForkRepositoryBasicDto struct {

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

func (o ForkRepositoryBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForkRepositoryBasicDto struct{}"
	}

	return strings.Join([]string{"ForkRepositoryBasicDto", string(data)}, " ")
}
