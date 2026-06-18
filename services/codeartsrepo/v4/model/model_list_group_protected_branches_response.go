package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupProtectedBranchesResponse Response Object
type ListGroupProtectedBranchesResponse struct {

	// **参数解释：** 代码组下保护分支列表。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Body           *[]GroupProtectedBranchApiDto `json:"body,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListGroupProtectedBranchesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupProtectedBranchesResponse struct{}"
	}

	return strings.Join([]string{"ListGroupProtectedBranchesResponse", string(data)}, " ")
}
