package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAdOuUsersRequest Request Object
type ListAdOuUsersRequest struct {

	// OU的域名地址。
	OuDn string `json:"ou_dn"`

	// 用户名，支持模糊查询。
	UserName *string `json:"user_name,omitempty"`

	// 用户是否已存在。
	HasExisted *bool `json:"has_existed,omitempty"`

	// 用于分页查询，返回桌面数量限制。如果不指定，则返回所有符合条件的桌面。
	Limit *int32 `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListAdOuUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAdOuUsersRequest struct{}"
	}

	return strings.Join([]string{"ListAdOuUsersRequest", string(data)}, " ")
}
