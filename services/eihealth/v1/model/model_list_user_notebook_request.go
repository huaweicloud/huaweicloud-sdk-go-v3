package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserNotebookRequest Request Object
type ListUserNotebookRequest struct {

	// notebook名称。
	NotebookName *string `json:"notebook_name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 镜像名称。
	ImageName *string `json:"image_name,omitempty"`

	// 是否仅展示当前用户创建的notebook。
	IsCreator *bool `json:"is_creator,omitempty"`

	// 空间名称列表。
	EihealthProjectNames *[]string `json:"eihealth_project_names,omitempty"`

	// 作业运行状态列表，支持RUNNING|ABNORMAL|HIBERNATE|SUCCEEDED|CREATING|DELETING|UPDATING|CREATEDFAILED|DELETEDFAILED|UPDATEDFAILED|UNKNOWN。
	Statuses *[]string `json:"statuses,omitempty"`

	// 最小创建时间。
	StartCreateTime *int64 `json:"start_create_time,omitempty"`

	// 最大创建时间。
	EndCreateTime *int64 `json:"end_create_time,omitempty"`

	// 最小更新时间。
	StartUpdateTime *int64 `json:"start_update_time,omitempty"`

	// 最大更新时间。
	EndUpdateTime *int64 `json:"end_update_time,omitempty"`

	// 排序规则，目前默认时间降序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序规则，默认更新时间降序，支持根据create_time|update_time排序。
	SortBy *string `json:"sort_by,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListUserNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserNotebookRequest struct{}"
	}

	return strings.Join([]string{"ListUserNotebookRequest", string(data)}, " ")
}
