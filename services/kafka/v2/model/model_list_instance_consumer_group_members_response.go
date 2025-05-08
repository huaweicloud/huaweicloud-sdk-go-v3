package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupMembersResponse Response Object
type ListInstanceConsumerGroupMembersResponse struct {

	// 成员详情
	Members *[]GroupMemberEntity `json:"members,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceConsumerGroupMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupMembersResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupMembersResponse", string(data)}, " ")
}
