package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TasksQueryInfo struct {

	// 测试任务URI集合
	Uris *[]string `json:"uris,omitempty"`

	// 关键字查询，任务名或编号
	Keyword *string `json:"keyword,omitempty"`

	// 标签集合
	Tags *[]string `json:"tags,omitempty"`

	// 是否是我的
	Own *bool `json:"own,omitempty"`

	// 服务类型
	ServiceType *int32 `json:"service_type,omitempty"`

	// 发布版本号集合
	ReleaseDevList *[]string `json:"release_dev_list,omitempty"`

	// 结果Code集合
	ResultCodes *[]string `json:"result_codes,omitempty"`

	// 状态Code集合
	StatusCodes *[]string `json:"status_codes,omitempty"`

	// 责任人ID集合
	OwnerIds *[]string `json:"owner_ids,omitempty"`

	// 执行者ID集合
	ExecutorIds *[]string `json:"executor_ids,omitempty"`

	// 创建者ID集合
	CreatorIds *[]string `json:"creator_ids,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式
	SortType *string `json:"sort_type,omitempty"`

	// 当前页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 是否轮询查询
	IsPollingQuery *bool `json:"is_polling_query,omitempty"`

	// 是否获取关联用例列表
	IsQueryAssociatedCaseList *bool `json:"is_query_associated_case_list,omitempty"`

	// 计划开始时间过滤起始时间戳
	PlanStartStartTimestamp *int64 `json:"plan_start_start_timestamp,omitempty"`

	// 计划开始时间过滤结束时间戳
	PlanStartEndTimestamp *int64 `json:"plan_start_end_timestamp,omitempty"`

	// 计划结束时间过滤起始时间戳
	PlanEndStartTimestamp *int64 `json:"plan_end_start_timestamp,omitempty"`

	// 计划结束时间过滤结束时间戳
	PlanEndEndTimestamp *int64 `json:"plan_end_end_timestamp,omitempty"`

	// 测试套超期状态过滤，超期状态值分别为：无状态(null)、未超期(0)、即将超期(1)、已超期(2)、延期完成(3)、按期完成(4)
	ExpirationStatusList *[]int32 `json:"expiration_status_list,omitempty"`
}

func (o TasksQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TasksQueryInfo struct{}"
	}

	return strings.Join([]string{"TasksQueryInfo", string(data)}, " ")
}
