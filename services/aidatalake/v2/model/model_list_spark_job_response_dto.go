package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSparkJobResponseDto 查看Spark作业的响应参数。
type ListSparkJobResponseDto struct {

	// **参数解释**：Spark作业ID，用于唯一标识该作业。 **取值范围**：采用UUID格式，长度为36个字符，例如：80ceaaff-3cfc-4162-a56f-70031ea4fa91。
	JobId string `json:"job_id"`

	// **参数解释**：Spark作业事务ID，用于防止重复提交。 **取值范围**：采用UUID格式，长度为36个字符，例如：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx。
	ClientToken *string `json:"client_token,omitempty"`

	// **参数解释**：Spark作业名称，用于标识作业。 **取值范围**：长度为1~128个字符。
	Name string `json:"name"`

	// **参数解释**：端点名称，用于指定Spark作业执行环境。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。
	EndpointName string `json:"endpoint_name"`

	// **参数解释**：作业类型，用于标识作业的类型。 **取值范围**： - spark_jar_job：Spark jar作业。 - spark_python_job：Python Spark作业。 - spark_sql_scripting_job：SQL脚本作业（预留类型）。
	JobType ListSparkJobResponseDtoJobType `json:"job_type"`

	// **参数解释**：Spark作业状态，用于标识作业当前执行状态。 **取值范围**： - PENDING：启动中。 - QUEUED：排队中。 - RUNNING：运行中。 - CANCELING：取消中。 - CANCELED：已取消。 - FAILED：运行失败。 - SUCCEED：运行成功。 - QUEUED_TIMEOUT：排队超时。 - RUNNING_TIMEOUT：运行超时。
	State string `json:"state"`

	// **参数解释**：作业状态消息，当作业异常结束时显示相关信息。 **取值范围**：长度为0~512个字符。
	StateMessage *string `json:"state_message,omitempty"`

	// **参数解释**：作业重试次数，用于记录作业失败后的重试次数。 **取值范围**：最小值为0。
	RetryTimes *int32 `json:"retry_times,omitempty"`

	// **参数解释**：作业创建时间，用于记录作业提交时间。 **取值范围**：unix时间戳，单位为毫秒。
	CreateTime int64 `json:"create_time"`

	// **参数解释**：作业开始运行时间，用于记录作业实际开始执行的时间。 **取值范围**：unix时间戳，单位为毫秒。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：作业结束时间，用于记录作业执行完成的时间。 **取值范围**：unix时间戳，单位为毫秒。
	EndTime *int64 `json:"end_time,omitempty"`

	CreateUser *SparkCreateUser `json:"create_user"`

	// **参数解释**：日志归档路径OBS URL，用于查看作业执行日志。 **取值范围**：OBS URL格式，长度为1~1024个字符。
	LogUrl *string `json:"log_url,omitempty"`
}

func (o ListSparkJobResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobResponseDto struct{}"
	}

	return strings.Join([]string{"ListSparkJobResponseDto", string(data)}, " ")
}

type ListSparkJobResponseDtoJobType struct {
	value string
}

type ListSparkJobResponseDtoJobTypeEnum struct {
	SPARK_JAR_JOB           ListSparkJobResponseDtoJobType
	SPARK_PYTHON_JOB        ListSparkJobResponseDtoJobType
	SPARK_SQL_SCRIPTING_JOB ListSparkJobResponseDtoJobType
}

func GetListSparkJobResponseDtoJobTypeEnum() ListSparkJobResponseDtoJobTypeEnum {
	return ListSparkJobResponseDtoJobTypeEnum{
		SPARK_JAR_JOB: ListSparkJobResponseDtoJobType{
			value: "spark_jar_job",
		},
		SPARK_PYTHON_JOB: ListSparkJobResponseDtoJobType{
			value: "spark_python_job",
		},
		SPARK_SQL_SCRIPTING_JOB: ListSparkJobResponseDtoJobType{
			value: "spark_sql_scripting_job",
		},
	}
}

func (c ListSparkJobResponseDtoJobType) Value() string {
	return c.value
}

func (c ListSparkJobResponseDtoJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSparkJobResponseDtoJobType) UnmarshalJSON(b []byte) error {
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
