package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryBriefDto struct {

	// **参数解释：** 仓库ID。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 仓库路径。 **取值范围：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 仓库完整路径。 **取值范围：** 不涉及。
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`

	// **参数解释：** 仓库所属项目名称。 **取值范围：** 不涉及。
	ProjectName *string `json:"project_name,omitempty"`
}

func (o RepositoryBriefDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryBriefDto struct{}"
	}

	return strings.Join([]string{"RepositoryBriefDto", string(data)}, " ")
}
