package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicRepositoryDto struct {

	// **参数解释：** 仓库ID **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库名称。 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 命名空间。 **约束限制：** 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释：** 仓库路径。 **约束限制：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 开发模式。 枚举值：normal：合并请求模式；CR：变更请求模式。 **约束限制：** 不涉及。
	DevelopMode *string `json:"develop_mode,omitempty"`

	// **参数解释：** 仓库可见性。 枚举值：private：私有；internal：内部公开；public：公开。 **约束限制：** 不涉及。
	Visibility *string `json:"visibility,omitempty"`

	// **参数解释：** 安全级别。 枚举值：product_internal：项目内公开；tenant_internal：租户内公开。 **约束限制：** 不涉及。
	Security *string `json:"security,omitempty"`

	// **参数解释：** 包含命名空间的仓库名称。 **约束限制：** 不涉及。
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// **参数解释：** 是否归档。true：已归档；false：未归档。 **约束限制：** 不涉及。
	Archived *bool `json:"archived,omitempty"`

	// **参数解释：** 仓库状态。 **约束限制：** 不涉及。
	Status *string `json:"status,omitempty"`
}

func (o BasicRepositoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicRepositoryDto struct{}"
	}

	return strings.Join([]string{"BasicRepositoryDto", string(data)}, " ")
}
