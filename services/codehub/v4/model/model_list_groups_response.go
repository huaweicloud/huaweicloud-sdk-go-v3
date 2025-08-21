package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsResponse Response Object
type ListGroupsResponse struct {

	// **参数解释：** 代码组列表。
	Body           *[]GroupBaseDto `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupsResponse", string(data)}, " ")
}
