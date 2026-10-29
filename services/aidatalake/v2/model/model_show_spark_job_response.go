package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSparkJobResponse Response Object
type ShowSparkJobResponse struct {

	// **参数解释**：Spark作业ID，用于唯一标识该作业。 **取值范围**：采用UUID格式，长度为36个字符，例如：80ceaaff-3cfc-4162-a56f-70031ea4fa91。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：Spark作业事务ID，用于防止重复提交。 **取值范围**：采用UUID格式，长度为36个字符，例如：xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx。
	ClientToken *string `json:"client_token,omitempty"`

	// **参数解释**：Spark作业名称，用于标识作业。 **取值范围**：长度为1~128个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**：端点名称，用于指定Spark作业执行环境。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：Catalog名称，用于指定作业使用的数据目录。 **取值范围**：长度为1~128个字符。
	CatalogName *string `json:"catalog_name,omitempty"`

	// **参数解释**：自定义委托的委托名，用于作业操作OBS对象、转储日志、访问DLI元数据等。 **取值范围**：长度为1~64个字符。
	JobAgency *string `json:"job_agency,omitempty"`

	// **参数解释**：Spark作业状态，用于标识作业当前执行状态。 **取值范围**： - PENDING：启动中。 - QUEUED：排队中。 - RUNNING：运行中。 - CANCELING：取消中。 - CANCELED：已取消。 - FAILED：运行失败。 - QUEUED_TIMEOUT：排队超时。 - RUNNING_TIMEOUT：运行超时。 - SUCCEED：运行成功。
	State *string `json:"state,omitempty"`

	// **参数解释**：作业状态消息，当作业异常结束时显示相关信息。 **取值范围**：长度为0~512个字符。
	StateMessage *string `json:"state_message,omitempty"`

	JobConfig *ShowSparkJobArtifactResponse `json:"job_config,omitempty"`

	ResourceConfig *ShowSparkResourceConfigResponse `json:"resource_config,omitempty"`

	RestoreStrategy *ShowSparkRestoreStrategyResponse `json:"restore_strategy,omitempty"`

	// **参数解释**：用户自定义Spark参数配置。 **取值范围**：长度为0~1024个字符。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	MetricStatistics *SparkJobMetricStatistics `json:"metric_statistics,omitempty"`

	LoggingConfig *ShowSparkLoggingConfigResponse `json:"logging_config,omitempty"`

	Image *ShowSparkJobImageConfigResponse `json:"image,omitempty"`

	// **参数解释**：作业重试次数，用于记录作业失败后的重试次数。 **取值范围**：最小值为0。
	RetryTimes *int32 `json:"retry_times,omitempty"`

	// **参数解释**：作业标签列表，用于标识和分类作业。
	Labels *[]ShowSparkJobLabelResponse `json:"labels,omitempty"`

	// **参数解释**：作业创建时间，用于记录作业提交时间。 **取值范围**：unix时间戳，单位为毫秒。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**：作业开始运行时间，用于记录作业实际开始执行的时间。 **取值范围**：unix时间戳，单位为毫秒。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：作业结束时间，用于记录作业执行完成的时间。 **取值范围**：unix时间戳，单位为毫秒。
	EndTime *int64 `json:"end_time,omitempty"`

	CreateUser *SparkCreateUser `json:"create_user,omitempty"`

	// **参数解释**：日志归档路径OBS URL，用于查看作业执行日志。 **取值范围**：OBS URL格式，长度为1~1024个字符。
	LogUrl *string `json:"log_url,omitempty"`

	// **参数解释**：Spark作业描述信息，用于说明作业用途。 **取值范围**：长度为0~512个字符。
	Description *string `json:"description,omitempty"`

	SparkUi        *SparkUiSharedInfo `json:"spark_ui,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowSparkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSparkJobResponse", string(data)}, " ")
}
