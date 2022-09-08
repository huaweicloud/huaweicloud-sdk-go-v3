package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAvailableRdsListResponse struct {

	// 获取创建逻辑库可用数据库实例信息列表的集合。
	Instances *[]QueryAvailableRdsList `json:"instances,omitempty"`

	// 分页参数: 起始值。
	Offset *int32 `json:"offset,omitempty"`

	// 分页参数：每页多少条。
	Limit *int32 `json:"limit,omitempty"`

	// 集合总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAvailableRdsListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableRdsListResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableRdsListResponse", string(data)}, " ")
}
