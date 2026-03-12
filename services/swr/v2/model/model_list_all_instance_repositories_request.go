package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllInstanceRepositoriesRequest Request Object
type ListAllInstanceRepositoriesRequest struct {

	// 分页查询时的查询标记，使用上一次接口调用返回的next_marker值，默认值从第一条数据查询。**注意：marker和limit参数需要配套使用。**
	Marker *string `json:"marker,omitempty"`

	// 条目数量，用于分页查询，默认值为100，最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 仓库名称
	Name *string `json:"name,omitempty"`
}

func (o ListAllInstanceRepositoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllInstanceRepositoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAllInstanceRepositoriesRequest", string(data)}, " ")
}
