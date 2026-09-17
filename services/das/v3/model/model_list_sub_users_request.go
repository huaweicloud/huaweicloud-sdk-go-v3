package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubUsersRequest Request Object
type ListSubUsersRequest struct {

	// 搜索关键字
	Keywords *string `json:"keywords,omitempty"`

	// 连接ID
	ConnectionId *string `json:"connection_id,omitempty"`
}

func (o ListSubUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubUsersRequest struct{}"
	}

	return strings.Join([]string{"ListSubUsersRequest", string(data)}, " ")
}
