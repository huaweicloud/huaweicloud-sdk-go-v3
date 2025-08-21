package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryProtectedTagBodyDto struct {

	// **参数解释：** 保护Tag名或通配符列表。 **约束限制：** 必传，每项需要至少能匹配到一个Tag。 **取值范围：** 字符串 **默认取值：** 不涉及
	Names []string `json:"names"`

	// **参数解释：** 事件设置列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Actions *[]RepositoryProtectedTagActionBodyDto `json:"actions,omitempty"`
}

func (o RepositoryProtectedTagBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryProtectedTagBodyDto struct{}"
	}

	return strings.Join([]string{"RepositoryProtectedTagBodyDto", string(data)}, " ")
}
