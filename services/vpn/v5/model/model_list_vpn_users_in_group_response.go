package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnUsersInGroupResponse Response Object
type ListVpnUsersInGroupResponse struct {

	// 用户列表信息
	Users *[]VpnUserInGroup `json:"users,omitempty"`

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ListVpnUsersInGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnUsersInGroupResponse struct{}"
	}

	return strings.Join([]string{"ListVpnUsersInGroupResponse", string(data)}, " ")
}
