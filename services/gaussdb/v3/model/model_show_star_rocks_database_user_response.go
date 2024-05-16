package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStarRocksDatabaseUserResponse Response Object
type ShowStarRocksDatabaseUserResponse struct {

	// 数据库账户信息。
	UserDetails *[]ShowStarRocksDatabaseUsersUserDetails `json:"user_details,omitempty"`

	// 数据库账户数量。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowStarRocksDatabaseUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStarRocksDatabaseUserResponse struct{}"
	}

	return strings.Join([]string{"ShowStarRocksDatabaseUserResponse", string(data)}, " ")
}
