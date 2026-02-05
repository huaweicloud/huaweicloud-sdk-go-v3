package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNonMigrationUsersResponse Response Object
type ListNonMigrationUsersResponse struct {

	// 热点时不迁移用户id总数量。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 热点时不迁移用户id列表。
	NonMigrateUsers *[]UserInfo `json:"non_migrate_users,omitempty"`
	HttpStatusCode  int         `json:"-"`
}

func (o ListNonMigrationUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNonMigrationUsersResponse struct{}"
	}

	return strings.Join([]string{"ListNonMigrationUsersResponse", string(data)}, " ")
}
