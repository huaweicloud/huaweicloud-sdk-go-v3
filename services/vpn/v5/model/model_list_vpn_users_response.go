package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnUsersResponse Response Object
type ListVpnUsersResponse struct {

	// 用户列表信息
	Users *[]VpnUser `json:"users,omitempty"`

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVpnUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnUsersResponse struct{}"
	}

	return strings.Join([]string{"ListVpnUsersResponse", string(data)}, " ")
}
