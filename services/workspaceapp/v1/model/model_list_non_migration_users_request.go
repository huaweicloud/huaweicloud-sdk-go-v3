package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNonMigrationUsersRequest Request Object
type ListNonMigrationUsersRequest struct {

	// 热点会话迁移配置ID。
	ConfigId string `json:"config_id"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListNonMigrationUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNonMigrationUsersRequest struct{}"
	}

	return strings.Join([]string{"ListNonMigrationUsersRequest", string(data)}, " ")
}
