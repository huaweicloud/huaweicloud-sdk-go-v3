package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetIdentityStoreSummaryResponse Response Object
type GetIdentityStoreSummaryResponse struct {

	// 已创建的用户数量
	Users *int64 `json:"users,omitempty"`

	// 用户配额
	UsersQuota *int64 `json:"users_quota,omitempty"`

	// 已创建的用户组数量
	Groups *int64 `json:"groups,omitempty"`

	// 用户组配额
	GroupsQuota    *int64 `json:"groups_quota,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o GetIdentityStoreSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetIdentityStoreSummaryResponse struct{}"
	}

	return strings.Join([]string{"GetIdentityStoreSummaryResponse", string(data)}, " ")
}
