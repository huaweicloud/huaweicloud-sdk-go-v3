package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRunsRequest Request Object
type ListRunsRequest struct {

	// 当前偏移量，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的最大作业个数，范围: [1, 100]。默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 用于查询开始时间在该时间点之后的作业。时间格式为ISO日期时间格式yyyy-MM-dd'T'HH:mm:ss.SSS。
	StartTime *string `json:"start_time,omitempty"`

	// 用于查询开始时间在该时间点之前的作业。时间格式为ISO日期时间格式yyyy-MM-dd'T'HH:mm:ss.SSS。
	EndTime *string `json:"end_time,omitempty"`

	// 仅当作业类型为SqlJob时可用。指定sql片段作为作业过滤条件，不区分大小写。
	SqlPattern *string `json:"sql_pattern,omitempty"`

	// 仅当作业类型为SqlJob时可用。SQL作业类型。DDL, DCL, IMPORT, EXPORT, QUERY, INSERT, SELECT, DATA_MIGRATION, ANALYZE, OBS_SELECT, COMPLEX
	SqlType *string `json:"sql_type,omitempty"`

	// 作业类型。目前仅支持SqlJob
	JobType *string `json:"job_type,omitempty"`

	// 此作业的当前状态，包含提交（LAUNCHING）、运行中（RUNNING）、完成（FINISHED）、失败（FAILED）、取消（CANCELLED）
	Status *string `json:"status,omitempty"`

	// 指定作业排序字段，默认为从created_time（作业提交时间），支持duration（作业运行时长）、created_time（作业提交时间） 、job_name（作业名称）三种排序字段。
	OrderBy *string `json:"order_by,omitempty"`

	// 指定作业排序的升降序，默认为desc（降序），支持asc（升序）、desc（降序）两种排序方式。
	Order *ListRunsRequestOrder `json:"order,omitempty"`

	// 作业名称
	JobName *string `json:"job_name,omitempty"`
}

func (o ListRunsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRunsRequest struct{}"
	}

	return strings.Join([]string{"ListRunsRequest", string(data)}, " ")
}

type ListRunsRequestOrder struct {
	value string
}

type ListRunsRequestOrderEnum struct {
	ASC  ListRunsRequestOrder
	DESC ListRunsRequestOrder
}

func GetListRunsRequestOrderEnum() ListRunsRequestOrderEnum {
	return ListRunsRequestOrderEnum{
		ASC: ListRunsRequestOrder{
			value: "asc",
		},
		DESC: ListRunsRequestOrder{
			value: "desc",
		},
	}
}

func (c ListRunsRequestOrder) Value() string {
	return c.value
}

func (c ListRunsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRunsRequestOrder) UnmarshalJSON(b []byte) error {
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
