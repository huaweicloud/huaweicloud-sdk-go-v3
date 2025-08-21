package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupAddableMembersResponse Response Object
type ListGroupAddableMembersResponse struct {

	// **参数解释：** 成员列表。
	Body           *[]GroupBatchAddMemberDto `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListGroupAddableMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupAddableMembersResponse struct{}"
	}

	return strings.Join([]string{"ListGroupAddableMembersResponse", string(data)}, " ")
}
