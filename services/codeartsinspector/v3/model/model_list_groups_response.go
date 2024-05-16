package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupsResponse Response Object
type ListGroupsResponse struct {

	// 群组总数
	Total *int32 `json:"total,omitempty"`

	// 群组列表
	Items          *[]Group `json:"items,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupsResponse", string(data)}, " ")
}
