package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserDrugJobRequest Request Object
type ListUserDrugJobRequest struct {

	// 是否仅展示本人创建资源。
	IsCreator *bool `json:"is_creator,omitempty"`

	// 作业名称。
	JobName *string `json:"job_name,omitempty"`

	// 空间名称列表，多个值之间使用逗号分隔。
	EihealthProjectNames *[]string `json:"eihealth_project_names,omitempty"`

	// 标签列表。
	Labels *[]string `json:"labels,omitempty"`

	// 作业运行状态列表，支持WAITING|RUNNING|FINISHED|FAILED|CANCELLED。
	StatusList *[]string `json:"status_list,omitempty"`

	// 作业类型列表，支持DOCKING|OPTIMIZATION|ADMET|POC_MOL_DESIGN|SEARCH|GENERATION|CPI|FEP|POCKET_DETECTION|TARGET_OPT|CLUSTERING。
	Types *[]string `json:"types,omitempty"`

	// 最小创建时间。
	CreateStartTime *int64 `json:"create_start_time,omitempty"`

	// 最大创建时间。
	CreateEndTime *int64 `json:"create_end_time,omitempty"`

	// 最小结束时间。
	FinishStartTime *int64 `json:"finish_start_time,omitempty"`

	// 最大结束时间。
	FinishEndTime *int64 `json:"finish_end_time,omitempty"`

	// 排序规则，目前默认时间降序。
	SortDir *string `json:"sort_dir,omitempty"`

	// 排序规则，目前默认时间降序，支持根据create_time|finish_time|running_time|total_time排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 限制量，单次查询总量，必须由数字组成，默认为100，取值范围[1,1000]。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，查询起始偏移，必须由数字组成，默认为0，取值范围[0,100000000]。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListUserDrugJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserDrugJobRequest struct{}"
	}

	return strings.Join([]string{"ListUserDrugJobRequest", string(data)}, " ")
}
