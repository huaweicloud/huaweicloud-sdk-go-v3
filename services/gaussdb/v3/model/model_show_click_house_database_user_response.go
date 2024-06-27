package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClickHouseDatabaseUserResponse Response Object
type ShowClickHouseDatabaseUserResponse struct {

	// 数据库账户信息。
	UserDetails *[]ShowClickHouseDatabaseUsersUserDetails `json:"user_details,omitempty"`

	// 数据库账户数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowClickHouseDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClickHouseDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"ShowClickHouseDatabaseUserResponse", string(data)}, " ")
}
