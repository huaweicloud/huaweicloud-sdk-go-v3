package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlJobsRequest Request Object
type ListSqlJobsRequest struct {

	// 当前页码，默认为第一页。
	CurrentPage *int32 `json:"current-page,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	// 用于查询开始时间在该时间点之前的作业。时间格式为unix时间戳，单位：毫秒。
	End *int64 `json:"end,omitempty"`

	EngineType *string `json:"engine-type,omitempty"`

	JobStatus *string `json:"job-status,omitempty"`

	// 指定查询的作业类型，包含DDL、DCL、IMPORT、EXPORT、QUERY、INSERT，若要查询所有类型的作业，则传入ALL。
	JobType *ListSqlJobsRequestJobType `json:"job-type,omitempty"`

	// 指定作业排序方式，默认为start_time_desc（作业提交时间降序），支持duration_desc（作业运行时长降序）、duration_asc（作业运行时长升序）、start_time_desc（作业提交时间降序）、start_time_asc（作业提交时间升序）四种排序方式。
	Order *ListSqlJobsRequestOrder `json:"order,omitempty"`

	// 提交作业的用户名称
	Owner *string `json:"owner,omitempty"`

	// 每页显示的最大作业个数，范围: [1, 100]。默认值：50。
	PageSize *int32 `json:"page-size,omitempty"`

	// 指定queue_name作为作业过滤条件，查询在指定queue上运行的作业。
	QueueName *string `json:"queue_name,omitempty"`

	// 指定sql片段作为作业过滤条件，不区分大小写。
	SqlPattern *string `json:"sql_pattern,omitempty"`

	// 用于查询开始时间在该时间点之后的作业。时间格式为unix时间戳，单位：毫秒。
	Start *int64 `json:"start,omitempty"`

	TableName *string `json:"table_name,omitempty"`

	// 指定作业标签作为过滤条件，支持多标签过滤。格式为“key=value”，如：GET /v1.0/{project_id}/jobs?tags=k1%3Dv1，“=”需要转义为“%3D”，“k1”为标签键，“v1”为标签值。
	Tags *string `json:"tags,omitempty"`
}

func (o ListSqlJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobsRequest struct{}"
	}

	return strings.Join([]string{"ListSqlJobsRequest", string(data)}, " ")
}

type ListSqlJobsRequestJobType struct {
	value string
}

type ListSqlJobsRequestJobTypeEnum struct {
	ALL    ListSqlJobsRequestJobType
	DDL    ListSqlJobsRequestJobType
	DCL    ListSqlJobsRequestJobType
	IMPORT ListSqlJobsRequestJobType
	EXPORT ListSqlJobsRequestJobType
	QUERY  ListSqlJobsRequestJobType
	INSERT ListSqlJobsRequestJobType
}

func GetListSqlJobsRequestJobTypeEnum() ListSqlJobsRequestJobTypeEnum {
	return ListSqlJobsRequestJobTypeEnum{
		ALL: ListSqlJobsRequestJobType{
			value: "ALL",
		},
		DDL: ListSqlJobsRequestJobType{
			value: "DDL",
		},
		DCL: ListSqlJobsRequestJobType{
			value: "DCL",
		},
		IMPORT: ListSqlJobsRequestJobType{
			value: "IMPORT",
		},
		EXPORT: ListSqlJobsRequestJobType{
			value: "EXPORT",
		},
		QUERY: ListSqlJobsRequestJobType{
			value: "QUERY",
		},
		INSERT: ListSqlJobsRequestJobType{
			value: "INSERT",
		},
	}
}

func (c ListSqlJobsRequestJobType) Value() string {
	return c.value
}

func (c ListSqlJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlJobsRequestJobType) UnmarshalJSON(b []byte) error {
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

type ListSqlJobsRequestOrder struct {
	value string
}

type ListSqlJobsRequestOrderEnum struct {
	DURATION_DESC   ListSqlJobsRequestOrder
	DURATION_ASC    ListSqlJobsRequestOrder
	START_TIME_DESC ListSqlJobsRequestOrder
	START_TIME_ASC  ListSqlJobsRequestOrder
}

func GetListSqlJobsRequestOrderEnum() ListSqlJobsRequestOrderEnum {
	return ListSqlJobsRequestOrderEnum{
		DURATION_DESC: ListSqlJobsRequestOrder{
			value: "duration_desc",
		},
		DURATION_ASC: ListSqlJobsRequestOrder{
			value: "duration_asc",
		},
		START_TIME_DESC: ListSqlJobsRequestOrder{
			value: "start_time_desc",
		},
		START_TIME_ASC: ListSqlJobsRequestOrder{
			value: "start_time_asc",
		},
	}
}

func (c ListSqlJobsRequestOrder) Value() string {
	return c.value
}

func (c ListSqlJobsRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlJobsRequestOrder) UnmarshalJSON(b []byte) error {
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
