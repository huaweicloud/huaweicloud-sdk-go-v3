package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBackPoolsResponse struct {

	// 每页记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码
	Offset *int64 `json:"offset,omitempty"`

	// 当前查询条件的后向流量池总数
	Count *int64 `json:"count,omitempty"`

	// 当前页的后向流量池记录列表
	Pools          *[]BackPoolVo `json:"pools,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListBackPoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackPoolsResponse struct{}"
	}

	return strings.Join([]string{"ListBackPoolsResponse", string(data)}, " ")
}
