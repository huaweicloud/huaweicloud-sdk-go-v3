package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ForkedFromRepositorySimplestDto struct {

	// **参数解释：** 带命名空间的仓库名称。 **约束限制：** 不涉及。
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`

	// **参数解释：** 带命名空间的仓库路径。 **约束限制：** 不涉及。
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`

	// **参数解释：** 仓库名称。 **约束限制：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 命名空间。 **约束限制：** 不涉及。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释：** 仓库ID **约束限制：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 产品ID。 **约束限制：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 产品名称。 **约束限制：** 不涉及。
	ProjectName *string `json:"project_name,omitempty"`
}

func (o ForkedFromRepositorySimplestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ForkedFromRepositorySimplestDto struct{}"
	}

	return strings.Join([]string{"ForkedFromRepositorySimplestDto", string(data)}, " ")
}
