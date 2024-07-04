package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVhostsResponse Response Object
type ListVhostsResponse struct {

	// 当前显示数量
	Size *int32 `json:"size,omitempty"`

	// 查询结果总数
	Total *int32 `json:"total,omitempty"`

	// 查询的Vhost信息列表
	Items          *[]ShowVhostDetailResp `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListVhostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVhostsResponse struct{}"
	}

	return strings.Join([]string{"ListVhostsResponse", string(data)}, " ")
}
