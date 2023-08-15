package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSimPoolMembersResponse Response Object
type ListSimPoolMembersResponse struct {

	// 每页记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码
	Offset *int64 `json:"offset,omitempty"`

	// 当前流量池的成员总数
	Count *int64 `json:"count,omitempty"`

	// 查询出来的流量池成员记录列表
	PoolMembers    *[]PoolMemVo `json:"pool_members,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListSimPoolMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSimPoolMembersResponse struct{}"
	}

	return strings.Join([]string{"ListSimPoolMembersResponse", string(data)}, " ")
}
