package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScheduledTaskHistoryRequest Request Object
type ListScheduledTaskHistoryRequest struct {

	// ID of ScheduledTask
	TaskId string `json:"task_id"`

	// 工单ID
	Id *string `json:"id,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 开始时间参数的查询开始区间时间戳
	StartedStartTime *int64 `json:"started_start_time,omitempty"`

	// 开始时间参数的查询结束区间时间戳
	StartedEndTime *int64 `json:"started_end_time,omitempty"`

	// 结束时间参数的查询开始区间时间戳
	FinishedStartTime *int64 `json:"finished_start_time,omitempty"`

	// 结束时间参数的查询结束区间时间戳
	FinishedEndTime *int64 `json:"finished_end_time,omitempty"`

	// 上一页数据的最后一条记录的id
	Marker *string `json:"marker,omitempty"`

	// 分页指针
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit int32 `json:"limit"`

	// 排序字段名：支持 started_time,finished_time
	SortKey *ListScheduledTaskHistoryRequestSortKey `json:"sort_key,omitempty"`

	// 排序方式，asc升序，desc降序
	SortDir *ListScheduledTaskHistoryRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListScheduledTaskHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTaskHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTaskHistoryRequest", string(data)}, " ")
}

type ListScheduledTaskHistoryRequestSortKey struct {
	value string
}

type ListScheduledTaskHistoryRequestSortKeyEnum struct {
	STARTED_TIME  ListScheduledTaskHistoryRequestSortKey
	FINISHED_TIME ListScheduledTaskHistoryRequestSortKey
}

func GetListScheduledTaskHistoryRequestSortKeyEnum() ListScheduledTaskHistoryRequestSortKeyEnum {
	return ListScheduledTaskHistoryRequestSortKeyEnum{
		STARTED_TIME: ListScheduledTaskHistoryRequestSortKey{
			value: "started_time",
		},
		FINISHED_TIME: ListScheduledTaskHistoryRequestSortKey{
			value: "finished_time",
		},
	}
}

func (c ListScheduledTaskHistoryRequestSortKey) Value() string {
	return c.value
}

func (c ListScheduledTaskHistoryRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledTaskHistoryRequestSortKey) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListScheduledTaskHistoryRequestSortDir struct {
	value string
}

type ListScheduledTaskHistoryRequestSortDirEnum struct {
	ASC  ListScheduledTaskHistoryRequestSortDir
	DESC ListScheduledTaskHistoryRequestSortDir
}

func GetListScheduledTaskHistoryRequestSortDirEnum() ListScheduledTaskHistoryRequestSortDirEnum {
	return ListScheduledTaskHistoryRequestSortDirEnum{
		ASC: ListScheduledTaskHistoryRequestSortDir{
			value: "asc",
		},
		DESC: ListScheduledTaskHistoryRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListScheduledTaskHistoryRequestSortDir) Value() string {
	return c.value
}

func (c ListScheduledTaskHistoryRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledTaskHistoryRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
