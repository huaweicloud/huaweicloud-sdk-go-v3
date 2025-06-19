package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSqlJobsRequest Request Object
type ListSqlJobsRequest struct {

	// 参数解释:  当前页码，默认为第一页 示例: 55 约束限制:  无 取值范围: 大于1的整数 默认取值: 1
	CurrentPage *int32 `json:"current-page,omitempty"`

	// 参数解释:  数据库名称 示例: UQuery 约束限制:  长度小于等于16 取值范围: 无 默认取值: 无
	DbName *string `json:"db_name,omitempty"`

	// 参数解释:  用于查询开始时间在该时间点之前的作业。时间格式为unix时间戳，单位：毫秒 示例: 156789546456 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	End *int64 `json:"end,omitempty"`

	// 参数解释:  引擎类型。支持配置spark引擎或hetuEngine引擎 示例: spark 约束限制:  无 取值范围: spark、hetuEngine 默认取值: 无
	EngineType *ListSqlJobsRequestEngineType `json:"engine-type,omitempty"`

	// 参数解释:  指定查询的作业状态 示例: FINISHED 约束限制:  长度小于等于16 取值范围: LAUNCHING RUNNING FAILED CANCELLED COMPUTED SUSPENDED ACTIVE DELETED CREATING FINISHED SCALING 默认取值: 无
	JobStatus *ListSqlJobsRequestJobStatus `json:"job-status,omitempty"`

	// 参数解释:  作业id 示例: bac76d9b-2891-4c50-a920-b98d8634fd0f 约束限制:  无 取值范围: 无 默认取值: 无
	JobId *string `json:"job-id,omitempty"`

	// 参数解释:  指定查询的作业类型，包含DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE，若要查询所有类型的作业，则传入ALL。 示例: QUERY 约束限制:  无 取值范围: DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE、ALL 默认取值: 无
	JobType *ListSqlJobsRequestJobType `json:"job-type,omitempty"`

	// 参数解释:  指定作业排序方式 示例: duration_desc 约束限制:  无 取值范围: start_time_desc（作业提交时间降序） start_time_asc（作业提交时间升序） duration_desc（作业运行时长降序） duration_asc（作业运行时长升序） 默认取值: start_time_desc
	Order *ListSqlJobsRequestOrder `json:"order,omitempty"`

	// 参数解释:  提交作业的用户名称 示例: ei_dlics_d00352431 约束限制:  无 取值范围: 无 默认取值: 无
	Owner *string `json:"owner,omitempty"`

	// 参数解释:  每页显示的最大作业个数 示例: 5 约束限制:  无 取值范围: [1, 100] 默认取值: 50
	PageSize *int32 `json:"page-size,omitempty"`

	// 参数解释:  指定queue_name作为作业过滤条件，查询在指定queue上运行的作业 示例: q1 约束限制:  匹配正则表达式^(?!_)(?![0-9]+$)[A-Za-z0-9_]*$且长度在[1, 128]范围内的字符串 取值范围: 无 默认取值: 无
	QueueName *string `json:"queue_name,omitempty"`

	// 参数解释:  指定sql片段作为作业过滤条件，不区分大小写 示例: .*? 约束限制:  长度在[0, 1024]范围内的字符串 取值范围: 无 默认取值: 无
	SqlPattern *string `json:"sql_pattern,omitempty"`

	// 参数解释:  用于查询开始时间在该时间点之后的作业。时间格式为unix时间戳，单位：毫秒 示例: 156456784655 约束限制:  无 取值范围: 大于等于1的整数 默认取值: 无
	Start *int64 `json:"start,omitempty"`

	// 参数解释:  记录其操作的表名称。类型为Import和Export作业才有“table_name”属性 示例: CS_JOB 约束限制:  无 取值范围: 无 默认取值: 无
	TableName *string `json:"table_name,omitempty"`

	// 参数解释:  指定作业标签作为过滤条件，支持多标签过滤。格式为“key=value”，如：GET /v1.0/{project_id}/jobs?tags=k1%3Dv1，“=”需要转义为“%3D”，“k1”为标签键，“v1”为标签值 示例: key=value 约束限制:  格式为“key=value”的字符串 取值范围: 无 默认取值: 无
	Tags *string `json:"tags,omitempty"`

	// 参数解释:  指定查询的作业类型列表，包含DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE，若要查询所有类型的作业，则传入ALL。 示例: QUERY 约束限制:  无 取值范围: DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE、ALL 默认取值: 无
	JobTypes *[]ListSqlJobsRequestJobTypes `json:"job_types,omitempty"`
}

func (o ListSqlJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlJobsRequest struct{}"
	}

	return strings.Join([]string{"ListSqlJobsRequest", string(data)}, " ")
}

type ListSqlJobsRequestEngineType struct {
	value string
}

type ListSqlJobsRequestEngineTypeEnum struct {
	SPARK       ListSqlJobsRequestEngineType
	HETU_ENGINE ListSqlJobsRequestEngineType
}

func GetListSqlJobsRequestEngineTypeEnum() ListSqlJobsRequestEngineTypeEnum {
	return ListSqlJobsRequestEngineTypeEnum{
		SPARK: ListSqlJobsRequestEngineType{
			value: "spark",
		},
		HETU_ENGINE: ListSqlJobsRequestEngineType{
			value: "hetuEngine",
		},
	}
}

func (c ListSqlJobsRequestEngineType) Value() string {
	return c.value
}

func (c ListSqlJobsRequestEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlJobsRequestEngineType) UnmarshalJSON(b []byte) error {
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

type ListSqlJobsRequestJobStatus struct {
	value string
}

type ListSqlJobsRequestJobStatusEnum struct {
	LAUNCHING ListSqlJobsRequestJobStatus
	RUNNING   ListSqlJobsRequestJobStatus
	FAILED    ListSqlJobsRequestJobStatus
	CANCELLED ListSqlJobsRequestJobStatus
	COMPUTED  ListSqlJobsRequestJobStatus
	SUSPENDED ListSqlJobsRequestJobStatus
	ACTIVE    ListSqlJobsRequestJobStatus
	DELETED   ListSqlJobsRequestJobStatus
	CREATING  ListSqlJobsRequestJobStatus
	FINISHED  ListSqlJobsRequestJobStatus
	SCALING   ListSqlJobsRequestJobStatus
}

func GetListSqlJobsRequestJobStatusEnum() ListSqlJobsRequestJobStatusEnum {
	return ListSqlJobsRequestJobStatusEnum{
		LAUNCHING: ListSqlJobsRequestJobStatus{
			value: "LAUNCHING",
		},
		RUNNING: ListSqlJobsRequestJobStatus{
			value: "RUNNING",
		},
		FAILED: ListSqlJobsRequestJobStatus{
			value: "FAILED",
		},
		CANCELLED: ListSqlJobsRequestJobStatus{
			value: "CANCELLED",
		},
		COMPUTED: ListSqlJobsRequestJobStatus{
			value: "COMPUTED",
		},
		SUSPENDED: ListSqlJobsRequestJobStatus{
			value: "SUSPENDED",
		},
		ACTIVE: ListSqlJobsRequestJobStatus{
			value: "ACTIVE",
		},
		DELETED: ListSqlJobsRequestJobStatus{
			value: "DELETED",
		},
		CREATING: ListSqlJobsRequestJobStatus{
			value: "CREATING",
		},
		FINISHED: ListSqlJobsRequestJobStatus{
			value: "FINISHED",
		},
		SCALING: ListSqlJobsRequestJobStatus{
			value: "SCALING",
		},
	}
}

func (c ListSqlJobsRequestJobStatus) Value() string {
	return c.value
}

func (c ListSqlJobsRequestJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlJobsRequestJobStatus) UnmarshalJSON(b []byte) error {
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

type ListSqlJobsRequestJobType struct {
	value string
}

type ListSqlJobsRequestJobTypeEnum struct {
	DDL            ListSqlJobsRequestJobType
	DCL            ListSqlJobsRequestJobType
	IMPORT         ListSqlJobsRequestJobType
	EXPORT         ListSqlJobsRequestJobType
	QUERY          ListSqlJobsRequestJobType
	INSERT         ListSqlJobsRequestJobType
	DATA_MIGRATION ListSqlJobsRequestJobType
	UPDATE         ListSqlJobsRequestJobType
	DELETE         ListSqlJobsRequestJobType
	RESTART_QUEUE  ListSqlJobsRequestJobType
	SCALE_QUEUE    ListSqlJobsRequestJobType
	ALL            ListSqlJobsRequestJobType
}

func GetListSqlJobsRequestJobTypeEnum() ListSqlJobsRequestJobTypeEnum {
	return ListSqlJobsRequestJobTypeEnum{
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
		DATA_MIGRATION: ListSqlJobsRequestJobType{
			value: "DATA_MIGRATION",
		},
		UPDATE: ListSqlJobsRequestJobType{
			value: "UPDATE",
		},
		DELETE: ListSqlJobsRequestJobType{
			value: "DELETE",
		},
		RESTART_QUEUE: ListSqlJobsRequestJobType{
			value: "RESTART_QUEUE",
		},
		SCALE_QUEUE: ListSqlJobsRequestJobType{
			value: "SCALE_QUEUE",
		},
		ALL: ListSqlJobsRequestJobType{
			value: "ALL",
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

type ListSqlJobsRequestJobTypes struct {
	value string
}

type ListSqlJobsRequestJobTypesEnum struct {
	DDL    ListSqlJobsRequestJobTypes
	DCL    ListSqlJobsRequestJobTypes
	IMPORT ListSqlJobsRequestJobTypes
	EXPORT ListSqlJobsRequestJobTypes
	QUERY  ListSqlJobsRequestJobTypes
	INSERT ListSqlJobsRequestJobTypes
	UPDATE ListSqlJobsRequestJobTypes
	DELETE ListSqlJobsRequestJobTypes
}

func GetListSqlJobsRequestJobTypesEnum() ListSqlJobsRequestJobTypesEnum {
	return ListSqlJobsRequestJobTypesEnum{
		DDL: ListSqlJobsRequestJobTypes{
			value: "DDL",
		},
		DCL: ListSqlJobsRequestJobTypes{
			value: "DCL",
		},
		IMPORT: ListSqlJobsRequestJobTypes{
			value: "IMPORT",
		},
		EXPORT: ListSqlJobsRequestJobTypes{
			value: "EXPORT",
		},
		QUERY: ListSqlJobsRequestJobTypes{
			value: "QUERY",
		},
		INSERT: ListSqlJobsRequestJobTypes{
			value: "INSERT",
		},
		UPDATE: ListSqlJobsRequestJobTypes{
			value: "UPDATE",
		},
		DELETE: ListSqlJobsRequestJobTypes{
			value: "DELETE",
		},
	}
}

func (c ListSqlJobsRequestJobTypes) Value() string {
	return c.value
}

func (c ListSqlJobsRequestJobTypes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSqlJobsRequestJobTypes) UnmarshalJSON(b []byte) error {
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
