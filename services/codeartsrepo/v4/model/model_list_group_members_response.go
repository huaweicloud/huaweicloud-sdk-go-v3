package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMembersResponse Response Object
type ListGroupMembersResponse struct {

	// **参数解释：** 成员列表。
	Body           *[]GroupMemberDetailDto `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListGroupMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMembersResponse struct{}"
	}

	return strings.Join([]string{"ListGroupMembersResponse", string(data)}, " ")
}
