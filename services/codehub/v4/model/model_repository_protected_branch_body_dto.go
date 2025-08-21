package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryProtectedBranchBodyDto struct {

	// **参数解释：** 保护分支名或通配符数组。 **约束限制：** 必传，每项需要至少能匹配到一个分支。 **取值范围：** 字符串 **默认取值：** 不涉及
	Names []string `json:"names"`

	// **参数解释：** 事件设置列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Actions *[]ProtectedBranchProtectedActionBodyDto `json:"actions,omitempty"`
}

func (o RepositoryProtectedBranchBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryProtectedBranchBodyDto struct{}"
	}

	return strings.Join([]string{"RepositoryProtectedBranchBodyDto", string(data)}, " ")
}
