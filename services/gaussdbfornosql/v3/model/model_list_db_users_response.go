package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDbUsersResponse struct {

	// 数据库用户信息列表。
	Users *[]RedisDbUserInfo `json:"users,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListDbUsersResponse", string(data)}, " ")
}
