package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlJob 作业信息。
type SqlJob struct {

	// 参数解释:  作业ID 示例: 6d2146a0-c2d5-41bd-8ca0-ca9694ada992 约束限制:  无 取值范围: 无 默认取值: 无
	JobId string `json:"job_id"`

	// 参数解释:  指定查询的作业类型，包含DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE，若要查询所有类型的作业，则传入ALL。 示例: QUERY 约束限制:  无 取值范围: DDL、DCL、IMPORT、EXPORT、QUERY、INSERT、DATA_MIGRATION、UPDATE、DELETE、RESTART_QUEUE、SCALE_QUEUE、ALL 默认取值: 无
	JobType string `json:"job_type"`

	// 作业提交的对列 示例: ci_sql 约束限制:  无 取值范围: 无 默认取值: 无
	QueueName string `json:"queue_name"`

	// 提交作业的用户 示例: tenant1 约束限制:  无 取值范围: 无 默认取值: 无
	Owner string `json:"owner"`

	// 作业开始的时间。是单位为“毫秒”的时间戳 示例: 1502349803729 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	StartTime int64 `json:"start_time"`

	// 作业运行时长，单位毫秒 示例: 30000 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	Duration *int64 `json:"duration,omitempty"`

	// 此作业的当前状态 示例: CANCELLED 约束限制:  无 取值范围: LAUNCHING（提交中） RUNNING（运行中） FINISHED（完成） FAILED（失败） CANCELLED（取消） 默认取值: 无
	Status SqlJobStatus `json:"status"`

	// Insert作业执行过程中扫描的记录条数 示例: 66 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	InputRowCount *int64 `json:"input_row_count,omitempty"`

	// Insert作业执行过程中扫描到的错误记录数 示例: 666 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	BadRowCount *int64 `json:"bad_row_count,omitempty"`

	// 作业执行过程中扫描文件的大小 示例: 5 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	InputSize int64 `json:"input_size"`

	// 当前作业返回的结果总条数或insert作业插入的总条数 示例: 55 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	ResultCount int64 `json:"result_count"`

	// 记录其操作的表所在的数据库名称。类型为Import和Export作业才有“database_name”属性 示例: db2 约束限制:  无 取值范围: 无 默认取值: 无
	DatabaseName *string `json:"database_name,omitempty"`

	// 记录其操作的表名称。类型为Import和Export作业才有“table_name”属性 示例: t2 约束限制:  无 取值范围: 无 默认取值: 无
	TableName *string `json:"table_name,omitempty"`

	// Import类型的作业，记录其导入的数据是否包括列名 示例: true 约束限制:  无 取值范围: true, false 默认取值: 无
	WithColumnHeader *bool `json:"with_column_header,omitempty"`

	// SQL查询的相关列信息的Json字符串 示例: {\\\"type\\\":\\\"struct\\\",\\\"fields\\\":[{\\\"name\\\":\\\"name\\\",\\\"type\\\":\\\"string\\\",\\\"nullable\\\":true,\\\"metadata\\\":{}},{\\\"name\\\":\\\"age\\\",\\\"type\\\":\\\"integer\\\",\\\"nullable\\\":true,\\\"metadata\\\":{}}]} 约束限制:  符合Json格式的字符串 取值范围: 无 默认取值: 无
	Detail string `json:"detail"`

	// 引擎类型 示例: spark 约束限制:  无 取值范围: spark, hetuEngine 默认取值: 无
	EngineType *SqlJobEngineType `json:"engine_type,omitempty"`

	// 作业执行的SQL语句 示例: select * from t_json_002 约束限制:  无 取值范围: 无 默认取值: 无
	Statement string `json:"statement"`

	// 作业标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 系统提示信息 示例: Navicat Data Modeler enables you to build high-quality conceptual, logical and physical data models for a wide variety of audiences. Navicat 15 has added support for the system-wide dark mode. Creativity is intelligence having fun. The On Startup feature allows you to control what tabs appear when you launch Navicat. In the Objects tab, you can use the List List, Detail Detail and ER Diagram ER Diagram buttons to change the object view. If your Internet Service Provider (ISP) does not provide direct access to its server, Secure Tunneling Protocol (SSH) / HTTP is another solution. A man’s best friends are his ten fingers. Export Wizard allows you to export data from tables, collections, views, or query results to any available formats. Navicat 15 has added support for the system-wide dark mode. Secure Sockets Layer(SSL) is a protocol for transmitting private documents via the Internet. A man’s best friends are his ten fingers. Navicat Monitor is a safe, simple and agentless remote server monitoring tool that is packed with powerful features to make your monitoring effective as possible. The past has no power over the present moment. Such sessions are also susceptible to session hijacking, where a malicious user takes over your session once you have authenticated. A man is not old until regrets take the place of dreams. Secure SHell (SSH) is a program to log in into another computer over a network, execute commands on a remote server, and move files from one machine to another. Secure SHell (SSH) is a program to log in into another computer over a network, execute commands on a remote server, and move files from one machine to another. Champions keep playing until they get it right. All journeys have secret destinations of which the traveler is unaware. Flexible settings enable you to set up a custom key for comparison and synchronization. Navicat authorizes you to make connection to remote servers running on different platforms (i.e. Windows, macOS, Linux and UNIX), and supports PAM and GSSAPI authentication. To successfully establish a new connection to local/remote server - no matter via SSL, SSH or HTTP, set the database login information in the General tab. It can also manage cloud databases such as Amazon Redshift, Amazon RDS, Alibaba Cloud. Features in Navicat are sophisticated enough to provide professional developers for all their specific needs, yet easy to learn for users who are new to database server. After logged in the Navicat Cloud feature, the Navigation pane will be divided into Navicat Cloud and My Connections sections. 约束限制:  无 取值范围: 无 默认取值: 无
	Message *string `json:"message,omitempty"`

	// 作业结束的时间。是单位为“毫秒”的时间戳 示例: 1502349803729 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	EndTime *int64 `json:"end_time,omitempty"`

	// 作业的CPU累计使用量 示例: 15376 约束限制:  无 取值范围: 无 默认取值: 无
	CpuCost *string `json:"cpu_cost,omitempty"`

	// 作业的输出字节数 示例: 35 约束限制:  无 取值范围: 无 默认取值: 无
	OutputByte *string `json:"output_byte,omitempty"`

	// 作业结果的存储路径。
	ResultPath *string `json:"result_path,omitempty"`

	// 作业结果的存储格式，当前只支持csv
	ResultFormat *string `json:"result_format,omitempty"`

	// 作业执行计划的存储路径。
	ExecutionDetailsPath *string `json:"execution_details_path,omitempty"`
}

func (o SqlJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJob struct{}"
	}

	return strings.Join([]string{"SqlJob", string(data)}, " ")
}

type SqlJobStatus struct {
	value string
}

type SqlJobStatusEnum struct {
	LAUNCHING SqlJobStatus
	RUNNING   SqlJobStatus
	FINISHED  SqlJobStatus
	FAILED    SqlJobStatus
	CANCELLED SqlJobStatus
}

func GetSqlJobStatusEnum() SqlJobStatusEnum {
	return SqlJobStatusEnum{
		LAUNCHING: SqlJobStatus{
			value: "LAUNCHING",
		},
		RUNNING: SqlJobStatus{
			value: "RUNNING",
		},
		FINISHED: SqlJobStatus{
			value: "FINISHED",
		},
		FAILED: SqlJobStatus{
			value: "FAILED",
		},
		CANCELLED: SqlJobStatus{
			value: "CANCELLED",
		},
	}
}

func (c SqlJobStatus) Value() string {
	return c.value
}

func (c SqlJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobStatus) UnmarshalJSON(b []byte) error {
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

type SqlJobEngineType struct {
	value string
}

type SqlJobEngineTypeEnum struct {
	SPARK       SqlJobEngineType
	HETU_ENGINE SqlJobEngineType
}

func GetSqlJobEngineTypeEnum() SqlJobEngineTypeEnum {
	return SqlJobEngineTypeEnum{
		SPARK: SqlJobEngineType{
			value: "spark",
		},
		HETU_ENGINE: SqlJobEngineType{
			value: "hetuEngine",
		},
	}
}

func (c SqlJobEngineType) Value() string {
	return c.value
}

func (c SqlJobEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlJobEngineType) UnmarshalJSON(b []byte) error {
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
