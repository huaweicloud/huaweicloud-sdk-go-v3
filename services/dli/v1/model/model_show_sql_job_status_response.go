package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSqlJobStatusResponse Response Object
type ShowSqlJobStatusResponse struct {

	// 参数解释:  作业ID 示例: 6d2146a0-c2d5-41bd-8ca0-ca9694ada992 约束限制:  无 取值范围: 任意字符串 默认取值: 无
	JobId *string `json:"job_id,omitempty"`

	// 参数解释:  作业类型 示例: 指定查询的作业类型，包含DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE，若要查询所有类型的作业，则传入ALL。 约束限制:  无 取值范围: DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE、ALL 默认取值: 无
	JobType *ShowSqlJobStatusResponseJobType `json:"job_type,omitempty"`

	// 参数解释:  作业提交的队列 示例: dli_sql 约束限制:  无 取值范围: 无 默认取值: 无
	QueueName *string `json:"queue_name,omitempty"`

	// 参数解释:  提交作业的用户 示例: ei_dlics_d00352431 约束限制:  无 取值范围: 无 默认取值: 无
	Owner *string `json:"owner,omitempty"`

	// 参数解释:  作业开始的时间。是单位为“毫秒”的时间戳 示例: 1502349803729 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	StartTime *int64 `json:"start_time,omitempty"`

	// 参数解释:  作业运行时长，单位毫秒 示例: 1000 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	Duration *int64 `json:"duration,omitempty"`

	// 参数解释:  此作业的当前状态 示例: CANCELLED 约束限制:  无 取值范围: LAUNCHING（提交） RUNNING（运行中） FINISH（完成） FAILED（失败） CANCELLED（取消） 默认取值: 无
	Status *ShowSqlJobStatusResponseStatus `json:"status,omitempty"`

	// 参数解释:  Insert作业执行过程中扫描的记录条数 示例: 66 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	InputRowCount *int64 `json:"input_row_count,omitempty"`

	// 参数解释:  Insert作业执行过程中扫描到的错误记录数 示例: 666 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	BadRowCount *int64 `json:"bad_row_count,omitempty"`

	// 参数解释:  作业执行过程中扫描文件的大小 示例: 5 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	InputSize *int64 `json:"input_size,omitempty"`

	// 参数解释:  当前作业返回的结果总条数或insert作业插入的总条数 示例: 55 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	ResultCount *int32 `json:"result_count,omitempty"`

	// 参数解释:  记录其操作的表所在的数据库名称。类型为Import和Export作业才有“database_name”属性 示例: db2 约束限制:  无 取值范围: 无 默认取值: 无
	DatabaseName *string `json:"database_name,omitempty"`

	// 参数解释:  记录其操作的表名称。类型为Import和Export作业才有“table_name”属性 示例: t2 约束限制:  无 取值范围: 无 默认取值: 无
	TableName *string `json:"table_name,omitempty"`

	// 参数解释:  SQL查询的相关列信息的Json字符串 示例: {\\\"type\\\":\\\"struct\\\",\\\"fields\\\":[{\\\"name\\\":\\\"name\\\",\\\"type\\\":\\\"string\\\",\\\"nullable\\\":true,\\\"metadata\\\":{}},{\\\"name\\\":\\\"age\\\",\\\"type\\\":\\\"integer\\\",\\\"nullable\\\":true,\\\"metadata\\\":{}}]} 约束限制:  符合JSON格式 取值范围: 无 默认取值: 无
	Detail *string `json:"detail,omitempty"`

	// 参数解释:  SQL配置参数信息Json字符串 示例: {\\\"type\\\":\\\"struct\\\",\\\"fields\\\":[{\\\"name\\\":\\\"name\\\",\\\"type\\\":\\\"string\\\",\\\"nullable\\\":true,\\\"metadata\\\":{}},{\\\"name\\\":\\\"age\\\",\\\"type\\\":\\\"integer\\\",\\\"nullable\\\":true,\\\"metadata\\\":{}}]} 约束限制:  符合JSON格式 取值范围: 无 默认取值: 无
	UserConf *string `json:"user_conf,omitempty"`

	// 参数解释:  查询结果OBS路径 示例: obs://bucketName/jobs/result/011c99a26ae84a1bb963a75e7637d3fd/2023/11/10/229b23ad-5033-40ed-ad70-fda900f2f04e 约束限制:  符合OBS路径格式，obs://_* 取值范围: 无 默认取值: 无
	ResultPath *string `json:"result_path,omitempty"`

	// 参数解释:  查询作业执行计划OBS路径 示例: obs://bucketName/jobs/execution_details/011c99a26ae84a1bb963a75e7637d3fd/2023/11/10/229b23ad-5033-40ed-ad70-fda900f2f04e 约束限制:  符合OBS路径格式，obs://_* 取值范围: 无 默认取值: 无
	ExecutionDetailsPath *string `json:"execution_details_path,omitempty"`

	// 参数解释:  查询结果格式 示例: csv 约束限制:  无 取值范围: csv 默认取值: 无
	ResultFormat *string `json:"result_format,omitempty"`

	// 参数解释:  作业执行的SQL语句 示例: select * from t_json_002 约束限制:  符合SQL格式 取值范围: 无 默认取值: 无
	Statement *string `json:"statement,omitempty"`

	// 参数解释:  执行请求是否成功。true表示请求执行成功 示例: true 约束限制:  无 取值范围: true、false 默认取值: 无
	IsSuccess *bool `json:"is_success,omitempty"`

	// 参数解释:  系统提示信息，执行成功时，信息可能为空 示例: success 约束限制:  无 取值范围: 无 默认取值: 无
	Message *string `json:"message,omitempty"`

	// 参数解释:  作业执行方式 示例: async 约束限制:  无 取值范围: async（同步） sync（异步） 默认取值: 无
	JobMode *string `json:"job_mode,omitempty"`

	// 作业标签
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSqlJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSqlJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowSqlJobStatusResponse", string(data)}, " ")
}

type ShowSqlJobStatusResponseJobType struct {
	value string
}

type ShowSqlJobStatusResponseJobTypeEnum struct {
	DDL            ShowSqlJobStatusResponseJobType
	DCL            ShowSqlJobStatusResponseJobType
	IMPORT         ShowSqlJobStatusResponseJobType
	EXPORT         ShowSqlJobStatusResponseJobType
	QUERY          ShowSqlJobStatusResponseJobType
	INSERT         ShowSqlJobStatusResponseJobType
	DATA_MIGRATION ShowSqlJobStatusResponseJobType
	UPDATE         ShowSqlJobStatusResponseJobType
	DELETE         ShowSqlJobStatusResponseJobType
	RESTART_QUEUE  ShowSqlJobStatusResponseJobType
	SCALE_QUEUE    ShowSqlJobStatusResponseJobType
}

func GetShowSqlJobStatusResponseJobTypeEnum() ShowSqlJobStatusResponseJobTypeEnum {
	return ShowSqlJobStatusResponseJobTypeEnum{
		DDL: ShowSqlJobStatusResponseJobType{
			value: "DDL",
		},
		DCL: ShowSqlJobStatusResponseJobType{
			value: "DCL",
		},
		IMPORT: ShowSqlJobStatusResponseJobType{
			value: "IMPORT",
		},
		EXPORT: ShowSqlJobStatusResponseJobType{
			value: "EXPORT",
		},
		QUERY: ShowSqlJobStatusResponseJobType{
			value: "QUERY",
		},
		INSERT: ShowSqlJobStatusResponseJobType{
			value: "INSERT",
		},
		DATA_MIGRATION: ShowSqlJobStatusResponseJobType{
			value: "DATA_MIGRATION",
		},
		UPDATE: ShowSqlJobStatusResponseJobType{
			value: "UPDATE",
		},
		DELETE: ShowSqlJobStatusResponseJobType{
			value: "DELETE",
		},
		RESTART_QUEUE: ShowSqlJobStatusResponseJobType{
			value: "RESTART_QUEUE",
		},
		SCALE_QUEUE: ShowSqlJobStatusResponseJobType{
			value: "SCALE_QUEUE",
		},
	}
}

func (c ShowSqlJobStatusResponseJobType) Value() string {
	return c.value
}

func (c ShowSqlJobStatusResponseJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobStatusResponseJobType) UnmarshalJSON(b []byte) error {
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

type ShowSqlJobStatusResponseStatus struct {
	value string
}

type ShowSqlJobStatusResponseStatusEnum struct {
	LAUNCHING ShowSqlJobStatusResponseStatus
	RUNNING   ShowSqlJobStatusResponseStatus
	FINISHED  ShowSqlJobStatusResponseStatus
	FAILED    ShowSqlJobStatusResponseStatus
	CANCELLED ShowSqlJobStatusResponseStatus
}

func GetShowSqlJobStatusResponseStatusEnum() ShowSqlJobStatusResponseStatusEnum {
	return ShowSqlJobStatusResponseStatusEnum{
		LAUNCHING: ShowSqlJobStatusResponseStatus{
			value: "LAUNCHING",
		},
		RUNNING: ShowSqlJobStatusResponseStatus{
			value: "RUNNING",
		},
		FINISHED: ShowSqlJobStatusResponseStatus{
			value: "FINISHED",
		},
		FAILED: ShowSqlJobStatusResponseStatus{
			value: "FAILED",
		},
		CANCELLED: ShowSqlJobStatusResponseStatus{
			value: "CANCELLED",
		},
	}
}

func (c ShowSqlJobStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowSqlJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSqlJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
