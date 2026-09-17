package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubUsersResponse Response Object
type ListSubUsersResponse struct {

	// IAM用户信息
	Data *[]SubUserInfo `json:"data,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 账号ID
	PrimaryAccountId *string `json:"primary_account_id,omitempty"`

	// 账号名称
	PrimaryAccountName *string `json:"primary_account_name,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ListSubUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubUsersResponse struct{}"
	}

	return strings.Join([]string{"ListSubUsersResponse", string(data)}, " ")
}
