package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserImageRequest Request Object
type ListUserImageRequest struct {

	// 是否仅展示本人创建资源。
	IsCreator *bool `json:"is_creator,omitempty"`

	// 镜像ID，支持精确搜索。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称，支持模糊搜索。
	ImageName *string `json:"image_name,omitempty"`

	// 镜像描述，支持模糊搜索。
	Description *string `json:"description,omitempty"`

	// 空间名称列表，支持查询多个空间下的镜像。
	EihealthProjectNames *[]string `json:"eihealth_project_names,omitempty"`

	// 源空间名称，支持模糊搜索。
	SourceProjectName *string `json:"source_project_name,omitempty"`

	// 镜像类型列表。
	Types *[]string `json:"types,omitempty"`

	// 最小创建时间。
	StartCreateTime *int64 `json:"start_create_time,omitempty"`

	// 最大创建时间。
	EndCreateTime *int64 `json:"end_create_time,omitempty"`

	// 最小更新时间。
	StartUpdateTime *int64 `json:"start_update_time,omitempty"`

	// 最大更新时间。
	EndUpdateTime *int64 `json:"end_update_time,omitempty"`

	// 排序规则，目前默认时间降序，支持根据create_time|update_time排序。
	SortBy *string `json:"sort_by,omitempty"`

	// 排序规则，目前默认时间降序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListUserImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserImageRequest struct{}"
	}

	return strings.Join([]string{"ListUserImageRequest", string(data)}, " ")
}
