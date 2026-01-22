package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserResponse Response Object
type ListUserResponse struct {

	// **参数解释**： 用户列表。
	Users *[]AmqpUser `json:"users,omitempty"`

	// **参数解释**： 总用户个数。 **取值范围**： 不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserResponse struct{}"
	}

	return strings.Join([]string{"ListUserResponse", string(data)}, " ")
}
