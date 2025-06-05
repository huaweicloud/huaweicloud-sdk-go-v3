package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserJobRequest Request Object
type ListUserJobRequest struct {

	// 是否仅展示本人创建资源。
	IsCreator *bool `json:"is_creator,omitempty"`

	// 作业名称，取值范围：长度为[1,63]，以小写字母开头，允许出现中划线（-）、小写字母和数字，且必须以小写字母或数字结尾。
	JobName *string `json:"job_name,omitempty"`

	// 空间名称列表，多个值之间使用逗号分隔。
	EihealthProjectNames *[]string `json:"eihealth_project_names,omitempty"`

	// 作业类型列表，支持workflow|app。
	Types *[]string `json:"types,omitempty"`

	// 作业运行状态列表，支持Succeeded|Running|Pending|Failed|Cancelling|Cancelled|Unknown。
	StatusList *[]string `json:"status_list,omitempty"`

	// 标签列表。
	Labels *[]string `json:"labels,omitempty"`

	// 开始创建时间。
	StartCreateTime *int64 `json:"start_create_time,omitempty"`

	// 结束创建时间。
	EndCreateTime *int64 `json:"end_create_time,omitempty"`

	// 开始完成时间。
	StartFinishTime *int64 `json:"start_finish_time,omitempty"`

	// 结束完成时间。
	EndFinishTime *int64 `json:"end_finish_time,omitempty"`

	// 排序规则，默认根据创建时间降序，支持create_time|finish_time。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序规则，目前默认时间降序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListUserJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserJobRequest struct{}"
	}

	return strings.Join([]string{"ListUserJobRequest", string(data)}, " ")
}
