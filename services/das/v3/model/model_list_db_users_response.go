package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDbUsersResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 数据库用户列表
	DbUsers        *[]DbUser `json:"db_users,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListDbUsersResponse", string(data)}, " ")
}
