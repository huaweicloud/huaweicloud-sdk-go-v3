package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListFlinkJobsRequest Request Object
type ListFlinkJobsRequest struct {

	// **参数解释**：工作空间ID，可通过[获取工作空间ID](aidatalake_02_0044.xml)获取。 **约束限制**：不涉及。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：上一页中最后一条记录id，查询第一页时传空值。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：指定每一页返回的最大条目数。 **约束限制**：不涉及。 **取值范围**：1~100。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页查询方向，用于控制翻页方向。 **约束限制**：与marker参数配合使用。当page_reverse=true时，必须传入marker参数，否则返回400错误。 **取值范围**：   - true：向前翻页（查询marker之前的数据）。   - false：向后翻页（查询marker之后的数据）。 **默认取值**：false。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之后的作业。unix时间戳，单位：毫秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTimeAfter *int64 `json:"create_time_after,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之前的作业。unix时间戳，单位：毫秒。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CreateTimeBefore *int64 `json:"create_time_before,omitempty"`

	// **参数解释**：作业运行的端点ID，用于按端点筛选作业。 **约束限制**：不涉及。 **取值范围**：采用UUID格式，固定36个字符。 **默认取值**：不涉及。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：通过端点名称检索的参数。 **约束限制**：不涉及。 **取值范围**：长度为1~64的中文、英文字母、数字、下划线、中划线、点号、空格的组合。 **默认取值**：不涉及。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：作业名称，用于按作业名称模糊筛选。 **约束限制**：不涉及。 **取值范围**：长度为1~128个字符，以字母或数字开头和结尾，中间可包含字母、数字、下划线和中划线。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：Flink作业ID，用于按作业ID筛选，可通过[创建Flink作业](RunFlinkJob.xml)接口返回值获取。 **约束限制**：不涉及。 **取值范围**：采用UUIDv7格式，固定36个字符。 **默认取值**：不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：已部署作业ID，用于按部署ID筛选作业。 **约束限制**：不涉及。 **取值范围**：长度为1~64个字符，只能由英文字母、数字、下划线和中划线组成。 **默认取值**：不涉及。
	DeploymentId *string `json:"deployment_id,omitempty"`

	// **参数解释**：Flink作业状态，用于按状态筛选作业列表。数组元素为FlinkJobState枚举，取值如下：   - CREATED：已创建。   - STARTING：启动中。   - RUNNING：运行中。   - CANCELLING：停止中。   - FAILED：已失败。   - CANCELLED：已停止。   - FINISHED：已完成。 **约束限制**：不涉及。
	State *[]FlinkJobState `json:"state,omitempty"`

	// **参数解释**：作业类型，用于按作业类型筛选列表。 **约束限制**：不涉及。 **取值范围**：   - FLINK_SQL_JOB：SQL作业。   - FLINK_JAR_JOB：JAR作业。   - FLINK_SQL_GATEWAY：SQL Gateway作业。   - FLINK_PYTHON_JOB：Python作业。 **默认取值**：不涉及。
	JobType *ListFlinkJobsRequestJobType `json:"job_type,omitempty"`

	// **参数解释**：执行模式，用于按执行模式筛选作业列表。 **约束限制**：不涉及。 **取值范围**：   - STREAMING：流模式。   - BATCH：批模式。 **默认取值**：不涉及。
	ExecutionMode *ListFlinkJobsRequestExecutionMode `json:"execution_mode,omitempty"`

	// **参数解释**：创建者用户ID（即IAM用户ID），用于按创建者筛选作业，获取方法请参见[获取账号ID](aidatalake_02_0030.xml)。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符，只能由英文字母和数字组成。 **默认取值**：不涉及。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释**：标签过滤条件，用于按标签筛选作业。格式为key=value，支持多个标签。 **约束限制**：最多10个标签。
	Labels *[]string `json:"labels,omitempty"`
}

func (o ListFlinkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlinkJobsRequest struct{}"
	}

	return strings.Join([]string{"ListFlinkJobsRequest", string(data)}, " ")
}

type ListFlinkJobsRequestJobType struct {
	value string
}

type ListFlinkJobsRequestJobTypeEnum struct {
	FLINK_SQL_JOB     ListFlinkJobsRequestJobType
	FLINK_JAR_JOB     ListFlinkJobsRequestJobType
	FLINK_SQL_GATEWAY ListFlinkJobsRequestJobType
	FLINK_PYTHON_JOB  ListFlinkJobsRequestJobType
}

func GetListFlinkJobsRequestJobTypeEnum() ListFlinkJobsRequestJobTypeEnum {
	return ListFlinkJobsRequestJobTypeEnum{
		FLINK_SQL_JOB: ListFlinkJobsRequestJobType{
			value: "FLINK_SQL_JOB",
		},
		FLINK_JAR_JOB: ListFlinkJobsRequestJobType{
			value: "FLINK_JAR_JOB",
		},
		FLINK_SQL_GATEWAY: ListFlinkJobsRequestJobType{
			value: "FLINK_SQL_GATEWAY",
		},
		FLINK_PYTHON_JOB: ListFlinkJobsRequestJobType{
			value: "FLINK_PYTHON_JOB",
		},
	}
}

func (c ListFlinkJobsRequestJobType) Value() string {
	return c.value
}

func (c ListFlinkJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlinkJobsRequestJobType) UnmarshalJSON(b []byte) error {
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

type ListFlinkJobsRequestExecutionMode struct {
	value string
}

type ListFlinkJobsRequestExecutionModeEnum struct {
	STREAMING ListFlinkJobsRequestExecutionMode
	BATCH     ListFlinkJobsRequestExecutionMode
}

func GetListFlinkJobsRequestExecutionModeEnum() ListFlinkJobsRequestExecutionModeEnum {
	return ListFlinkJobsRequestExecutionModeEnum{
		STREAMING: ListFlinkJobsRequestExecutionMode{
			value: "STREAMING",
		},
		BATCH: ListFlinkJobsRequestExecutionMode{
			value: "BATCH",
		},
	}
}

func (c ListFlinkJobsRequestExecutionMode) Value() string {
	return c.value
}

func (c ListFlinkJobsRequestExecutionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlinkJobsRequestExecutionMode) UnmarshalJSON(b []byte) error {
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
