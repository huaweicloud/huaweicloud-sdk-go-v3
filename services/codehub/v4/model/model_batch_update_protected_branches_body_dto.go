package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateProtectedBranchesBodyDto struct {

	// **参数解释：** 保护分支名列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Names []string `json:"names"`

	// **参数解释：** 保护事件列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Actions *[]ProtectedBranchProtectedActionBodyDto `json:"actions,omitempty"`
}

func (o BatchUpdateProtectedBranchesBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateProtectedBranchesBodyDto struct{}"
	}

	return strings.Join([]string{"BatchUpdateProtectedBranchesBodyDto", string(data)}, " ")
}
