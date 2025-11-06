package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryProtectedTagDto struct {

	// **参数解释：** 保护Tag ID。 **取值范围：** 1-2147483647
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 分支名称或通配符。 **取值范围：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 事件列表。 **取值范围：** 不涉及
	Actions *[]RepositoryProtectedActionDto `json:"actions,omitempty"`
}

func (o RepositoryProtectedTagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryProtectedTagDto struct{}"
	}

	return strings.Join([]string{"RepositoryProtectedTagDto", string(data)}, " ")
}
