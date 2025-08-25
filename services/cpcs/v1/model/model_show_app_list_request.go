package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppListRequest Request Object
type ShowAppListRequest struct {

	// 指定查询返回记录条数，默认值10
	PageSize *int32 `json:"page_size,omitempty"`

	// 索引位置，从page_num指定的下一条数据开始查询默认值为0
	PageNum *int32 `json:"page_num,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 所属的VPC名称
	VpcName *string `json:"vpc_name,omitempty"`

	// 排序属性，目前支持以下属性： - **create_time** : 应用的创建时间（默认）
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向，支持以下值： - **DESC** : 降序（默认） - **ASC** : 升序
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ShowAppListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppListRequest struct{}"
	}

	return strings.Join([]string{"ShowAppListRequest", string(data)}, " ")
}
