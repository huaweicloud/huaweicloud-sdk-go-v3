package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateProtectedTagsBodyDto struct {

	// **参数解释：** 保护Tag名或通配符列表。 **约束限制：** 必填 **取值范围：** 不涉及 **默认取值：** 不涉及。
	Names []string `json:"names"`

	// **参数解释：** 权限列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及。
	Actions *[]RepositoryProtectedTagActionBodyDto `json:"actions,omitempty"`
}

func (o BatchUpdateProtectedTagsBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateProtectedTagsBodyDto struct{}"
	}

	return strings.Join([]string{"BatchUpdateProtectedTagsBodyDto", string(data)}, " ")
}
