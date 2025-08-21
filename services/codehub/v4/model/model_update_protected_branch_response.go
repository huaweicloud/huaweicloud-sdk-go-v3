package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProtectedBranchResponse Response Object
type UpdateProtectedBranchResponse struct {

	// **参数解释：** 保护分支ID。 **取值范围：** 1-2147483647
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 分支名称或通配符。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 事件列表。 **取值范围：** 不涉及。
	Actions        *[]RepositoryProtectedActionDto `json:"actions,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o UpdateProtectedBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectedBranchResponse struct{}"
	}

	return strings.Join([]string{"UpdateProtectedBranchResponse", string(data)}, " ")
}
