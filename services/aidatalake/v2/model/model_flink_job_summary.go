package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJobSummary **参数解释**：Flink作业摘要信息（列表展示）。
type FlinkJobSummary struct {

	// **参数解释**：Flink作业ID。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	JobId string `json:"job_id"`

	// **参数解释**：作业名称。 **取值范围**：1~128个字符，以字母或数字开头和结尾，中间可包含字母、数字、下划线和中划线。
	Name string `json:"name"`

	// **参数解释**：作业运行的端点ID。 **取值范围**：采用UUID格式，固定36个字符。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：作业运行的端点名称。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：作业类型。 **取值范围**：   - FLINK_JAR_JOB：JAR作业。   - FLINK_SQL_JOB：SQL作业。   - FLINK_SQL_GATEWAY：SQL Gateway作业。   - FLINK_PYTHON_JOB：Python作业。
	JobType *FlinkJobSummaryJobType `json:"job_type,omitempty"`

	Status *FlinkJobStatus `json:"status"`

	ExecutionMode *FlinkJobExecutionMode `json:"execution_mode"`

	Metric *FlinkResourceMetric `json:"metric"`

	// **参数解释**：创建时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	CreateTime int64 `json:"create_time"`

	CreateUser *FlinkCreateUser `json:"create_user"`

	// **参数解释**：更新时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	UpdateTime int64 `json:"update_time"`

	UpdateUser *FlinkUpdateUser `json:"update_user"`

	// **参数解释**：日志归档路径OBS URL，用于查看作业执行日志。 **取值范围**：OBS URL格式，长度为1~1024个字符。
	LogUrl *string `json:"log_url,omitempty"`

	// **参数解释**：作业描述信息。 **取值范围**：1~512个字符。
	Description string `json:"description"`

	// **参数解释**：该作业是否开启失败告警通知。 **取值范围**：   - true：开启。作业运行失败或提交失败时会向客户发送告警通知。   - false：关闭。作业失败时不发送告警通知。 **默认取值**：false，不指定时默认关闭。
	AlertEnabled *bool `json:"alert_enabled,omitempty"`

	// **参数解释**：作业开始运行的时间戳（进入RUNNING状态的时间）。 **取值范围**：0 ~ 4102416000000，13位毫秒时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：作业结束时间戳（进入终态 FINISHED/CANCELLED/FAILED的时间）。 **取值范围**：0 ~ 4102416000000，13位毫秒时间戳。
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o FlinkJobSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobSummary struct{}"
	}

	return strings.Join([]string{"FlinkJobSummary", string(data)}, " ")
}

type FlinkJobSummaryJobType struct {
	value string
}

type FlinkJobSummaryJobTypeEnum struct {
	FLINK_JAR_JOB     FlinkJobSummaryJobType
	FLINK_SQL_JOB     FlinkJobSummaryJobType
	FLINK_SQL_GATEWAY FlinkJobSummaryJobType
	FLINK_PYTHON_JOB  FlinkJobSummaryJobType
}

func GetFlinkJobSummaryJobTypeEnum() FlinkJobSummaryJobTypeEnum {
	return FlinkJobSummaryJobTypeEnum{
		FLINK_JAR_JOB: FlinkJobSummaryJobType{
			value: "FLINK_JAR_JOB",
		},
		FLINK_SQL_JOB: FlinkJobSummaryJobType{
			value: "FLINK_SQL_JOB",
		},
		FLINK_SQL_GATEWAY: FlinkJobSummaryJobType{
			value: "FLINK_SQL_GATEWAY",
		},
		FLINK_PYTHON_JOB: FlinkJobSummaryJobType{
			value: "FLINK_PYTHON_JOB",
		},
	}
}

func (c FlinkJobSummaryJobType) Value() string {
	return c.value
}

func (c FlinkJobSummaryJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobSummaryJobType) UnmarshalJSON(b []byte) error {
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
