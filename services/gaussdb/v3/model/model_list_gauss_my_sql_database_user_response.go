package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListGaussMySqlDatabaseUserResponse struct {

	// 数据库用户列表。
	Users *[]ListGaussMySqlDatabaseUser `json:"users,omitempty"`

	// 实例的数据库用户总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGaussMySqlDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussMySqlDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlDatabaseUserResponse", string(data)}, " ")
}
