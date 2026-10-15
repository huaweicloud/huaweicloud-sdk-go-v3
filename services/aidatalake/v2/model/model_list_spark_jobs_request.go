package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSparkJobsRequest Request Object
type ListSparkJobsRequest struct {

	// **参数解释**：工作空间的ID。 **约束限制**：不涉及。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。 **默认取值**：不涉及。
	WorkspaceId string `json:"workspace_id"`

	// **参数解释**：作业ID游标位置，用于分页查询。 **约束限制**：首次查询可不传或传空字符串，后续查询传入上次返回的next_marker值。 **取值范围**：采用UUID格式，长度为36个字符，只能由英文大写字母、英文小写字母、数字及下划线和横线组成。 **默认取值**：不涉及。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：查询记录数，用于分页查询。 **约束限制**：不涉及。 **取值范围**：1~100之间的整数。 **默认取值**：10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向分页查询。默认为false，表示正向分页查询。 当设置为true时，表示查询上一页数据，需要配合marker参数使用。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：作业名称，用于模糊查询作业。 **约束限制**：不涉及。 **取值范围**：长度为0~128个字符。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之后的作业。 **约束限制**：不涉及。 **取值范围**：unix时间戳，单位为毫秒，最小值为1764061598000。 **默认取值**：不涉及。
	CreateTimeAfter *int64 `json:"create_time_after,omitempty"`

	// **参数解释**：用于查询创建时间在该时间点之前的作业。 **约束限制**：不涉及。 **取值范围**：unix时间戳，单位为毫秒，最小值为1764061598000。 **默认取值**：不涉及。
	CreateTimeBefore *int64 `json:"create_time_before,omitempty"`

	// **参数解释**：端点名称，用于指定Spark作业执行环境。 **约束限制**：不涉及。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。 **默认取值**：不涉及。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：Spark作业状态列表，用于过滤查询指定状态的作业。 **约束限制**：最多支持查询10种状态。 **取值范围**： - PENDING：启动中。 - QUEUED：排队中。 - RUNNING：运行中。 - CANCELING：取消中。 - CANCELED：已取消。 - FAILED：运行失败。 - SUCCEED：运行成功。 - QUEUED_TIMEOUT：排队超时。 - RUNNING_TIMEOUT：运行超时。
	States *[]string `json:"states,omitempty"`

	// **参数解释**：作业类型，用于过滤查询指定类型的作业。 **约束限制**：不涉及。 **取值范围**： - spark_jar_job：Spark jar作业。 - spark_python_job：Python Spark作业。 - spark_sql_scripting_job：SQL脚本作业（预留类型）。 **默认取值**：不涉及。
	JobType *ListSparkJobsRequestJobType `json:"job_type,omitempty"`

	// **参数解释**：Spark作业ID，用于精确查询指定作业。 **约束限制**：不涉及。 **取值范围**：只能由英文大写字母、英文小写字母、数字及下划线和横线组成，且长度为1~64个字符。 **默认取值**：不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：Spark作业创建者ID，用于过滤查询指定用户创建的作业。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。 **默认取值**：不涉及。
	CreateUserId *string `json:"create_user_id,omitempty"`

	// **参数解释**：Spark作业创建者名称,用于精确过滤查询指定用户创建的作业。 **约束限制**：不涉及。 **取值范围**：长度为1~256个字符。
	CreateUserName *string `json:"create_user_name,omitempty"`

	// **参数解释**：作业标签，用于按标签过滤查询作业。支持多标签过滤，格式为“key=value”，多个标签用逗号分隔。 **约束限制**：URL中“=”需要转义为“%3D”。例如：labels=k1%3Dv1，k2%3Dv2。 **取值范围**：长度为1~1024个字符。 **默认取值**：不涉及。
	Labels *string `json:"labels,omitempty"`
}

func (o ListSparkJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSparkJobsRequest struct{}"
	}

	return strings.Join([]string{"ListSparkJobsRequest", string(data)}, " ")
}

type ListSparkJobsRequestJobType struct {
	value string
}

type ListSparkJobsRequestJobTypeEnum struct {
	SPARK_JAR_JOB           ListSparkJobsRequestJobType
	SPARK_PYTHON_JOB        ListSparkJobsRequestJobType
	SPARK_SQL_SCRIPTING_JOB ListSparkJobsRequestJobType
}

func GetListSparkJobsRequestJobTypeEnum() ListSparkJobsRequestJobTypeEnum {
	return ListSparkJobsRequestJobTypeEnum{
		SPARK_JAR_JOB: ListSparkJobsRequestJobType{
			value: "spark_jar_job",
		},
		SPARK_PYTHON_JOB: ListSparkJobsRequestJobType{
			value: "spark_python_job",
		},
		SPARK_SQL_SCRIPTING_JOB: ListSparkJobsRequestJobType{
			value: "spark_sql_scripting_job",
		},
	}
}

func (c ListSparkJobsRequestJobType) Value() string {
	return c.value
}

func (c ListSparkJobsRequestJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSparkJobsRequestJobType) UnmarshalJSON(b []byte) error {
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
