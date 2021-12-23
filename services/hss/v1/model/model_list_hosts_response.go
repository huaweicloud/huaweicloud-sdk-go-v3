package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHostsResponse struct {
	// 总数

	TotalNum *int32 `json:"total_num,omitempty"`
	// 查询弹性云服务器状态列表

	DataList       *[]Host `json:"data_list,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsResponse struct{}"
	}

	return strings.Join([]string{"ListHostsResponse", string(data)}, " ")
}
