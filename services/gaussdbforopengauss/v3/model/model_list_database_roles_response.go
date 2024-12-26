package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseRolesResponse Response Object
type ListDatabaseRolesResponse struct {

	// 列表中每个元素表示一个用户/角色。
	Roles *[]GaussDbListDatabaseRoles `json:"roles,omitempty"`

	// 数据库用户/角色总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDatabaseRolesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseRolesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseRolesResponse", string(data)}, " ")
}
