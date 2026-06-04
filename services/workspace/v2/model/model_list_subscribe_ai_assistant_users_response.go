package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscribeAiAssistantUsersResponse Response Object
type ListSubscribeAiAssistantUsersResponse struct {

	// 订阅用户总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 订阅用户列表。
	Users *[]SubscribeUserBasicInfo `json:"users,omitempty"`

	// 订阅用户组列表。
	Usergroups *[]SubscribeUserGroupInfo `json:"usergroups,omitempty"`

	Project        *SubscribeAiAssistantListResponseProject `json:"project,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListSubscribeAiAssistantUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscribeAiAssistantUsersResponse struct{}"
	}

	return strings.Join([]string{"ListSubscribeAiAssistantUsersResponse", string(data)}, " ")
}
