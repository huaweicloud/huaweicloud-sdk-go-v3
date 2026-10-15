package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkJobResponse Response Object
type ShowFlinkJobResponse struct {

	// **参数解释**：Flink作业ID。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：客户端Token，用于接口调用幂等校验。 **取值范围**：长度为1~64个字符，只能由英文字母、数字、下划线和中划线组成。
	ClientToken *string `json:"client_token,omitempty"`

	// **参数解释**：作业名称。 **取值范围**：1~128个字符，以字母或数字开头和结尾，中间可包含字母、数字、下划线和中划线。
	Name *string `json:"name,omitempty"`

	// **参数解释**：已部署作业ID。 **取值范围**：1~64个字符，支持字母、数字、下划线和中划线。
	DeploymentId *string `json:"deployment_id,omitempty"`

	// **参数解释**：作业运行的端点ID。 **取值范围**：采用UUID格式，固定36个字符。
	EndpointId *string `json:"endpoint_id,omitempty"`

	// **参数解释**：作业运行的端点名称。 **取值范围**：只能由英文小写字母、数字及中划线组成，以英文小写字母开头，以英文小写字母或数字结尾，且长度为1~63个字符。
	EndpointName *string `json:"endpoint_name,omitempty"`

	Status *FlinkJobStatus `json:"status,omitempty"`

	// **参数解释**：用户自定义委托名称，用于授权作业访问OBS等外部资源，不指定时默认使用端点委托。 **取值范围**：1~64个字符。
	JobAgency *string `json:"job_agency,omitempty"`

	// **参数解释**：OBS桶名称，用于存储作业的Checkpoint、日志等数据。 **取值范围**：1~512个字符。
	ObsBucketName *string `json:"obs_bucket_name,omitempty"`

	JobConfig *ShowFlinkJobArtifactResponse `json:"job_config,omitempty"`

	ResourceConfig *ShowFlinkResourceConfigResponseBody `json:"resource_config,omitempty"`

	ImageConfig *ShowFlinkJobImageConfigResponseBody `json:"image_config,omitempty"`

	// **参数解释**：Flink配置参数，为键值对形式。key为Flink配置项名称（如taskmanager.numberOfTaskSlots）；value为对应配置值，长度1~1024个字符。除用户配置项外，还包含系统自动补充的保留参数（如execution.checkpointing.dir、execution.checkpointing.savepoint-dir、high-availability.storageDir，由系统根据作业配置自动生成）。
	FlinkConfig map[string]string `json:"flink_config,omitempty"`

	ExecutionMode *FlinkJobExecutionMode `json:"execution_mode,omitempty"`

	RestoreStrategy *ShowRestoreStrategyResponseBody `json:"restore_strategy,omitempty"`

	LoggingConfig *ShowFlinkLoggingConfigResponseBody `json:"logging_config,omitempty"`

	Metric *FlinkResourceMetric `json:"metric,omitempty"`

	// **参数解释**：创建时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	CreateUser *FlinkCreateUser `json:"create_user,omitempty"`

	// **参数解释**：更新时间戳。 **取值范围**：0~4102416000000，13位毫秒时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`

	UpdateUser *FlinkUpdateUser `json:"update_user,omitempty"`

	// **参数解释**：日志归档路径OBS URL，用于查看作业执行日志。 **取值范围**：OBS URL格式，长度为1~1024个字符。
	LogUrl *string `json:"log_url,omitempty"`

	// **参数解释**：标签，用于作业分类和筛选。
	Labels *[]FlinkLabel `json:"labels,omitempty"`

	// **参数解释**：描述信息。 **取值范围**：1~512个字符。
	Description *string `json:"description,omitempty"`

	// **参数解释**：作业开始运行的时间戳（进入RUNNING状态的时间）。 **取值范围**：0 ~ 4102416000000，13位毫秒时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：作业结束时间戳（进入终态 FINISHED/CANCELLED/FAILED的时间）。 **取值范围**：0 ~ 4102416000000，13位毫秒时间戳。
	EndTime        *int64 `json:"end_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"ShowFlinkJobResponse", string(data)}, " ")
}
