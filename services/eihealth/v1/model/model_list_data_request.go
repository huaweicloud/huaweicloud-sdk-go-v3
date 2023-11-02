package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataRequest Request Object
type ListDataRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]
	Offset *int32 `json:"offset,omitempty"`

	// 指定文件夹（项目名称:/路径）
	Path *string `json:"path,omitempty"`

	// 查询关键词
	SearchKey *string `json:"search_key,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为desc）
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序字段（支持type，create_time）
	SortKey *string `json:"sort_key,omitempty"`

	// 开始标签
	Marker *string `json:"marker,omitempty"`
}

func (o ListDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataRequest struct{}"
	}

	return strings.Join([]string{"ListDataRequest", string(data)}, " ")
}
