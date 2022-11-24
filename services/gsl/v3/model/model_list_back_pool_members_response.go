package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackPoolMembersResponse struct {

	// 每页记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码
	Offset *int64 `json:"offset,omitempty"`

	// 当前流量池的成员总数
	Count *int64 `json:"count,omitempty"`

	// 查询出来的流量池成员记录列表
	PoolMembers    *[]BackPoolMemVo `json:"pool_members,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListBackPoolMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackPoolMembersResponse struct{}"
	}

	return strings.Join([]string{"ListBackPoolMembersResponse", string(data)}, " ")
}
