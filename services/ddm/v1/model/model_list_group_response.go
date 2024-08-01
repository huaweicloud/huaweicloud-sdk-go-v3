package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupResponse Response Object
type ListGroupResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 组信息列表。
	GroupList      *[]GroupInfo `json:"group_list,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupResponse struct{}"
	}

	return strings.Join([]string{"ListGroupResponse", string(data)}, " ")
}
