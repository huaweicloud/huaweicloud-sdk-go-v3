package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubscribeAiAssistantUsersRequest Request Object
type ListSubscribeAiAssistantUsersRequest struct {

	// 用于分页查询，返回用户数量限制。如果不指定，则返回所有符合条件的用户。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始条数。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListSubscribeAiAssistantUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubscribeAiAssistantUsersRequest struct{}"
	}

	return strings.Join([]string{"ListSubscribeAiAssistantUsersRequest", string(data)}, " ")
}
