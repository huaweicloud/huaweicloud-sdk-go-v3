package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueueResponse Response Object
type ShowQueueResponse struct {

	// 请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 队列ID。
	QueueId *int64 `json:"queue_id,omitempty"`

	// 队列名称。说明：队列名称不区分大小写，系统会自动转换为小写。
	QueueName *string `json:"queueName,omitempty"`

	// 队列描述信息。
	Description *string `json:"description,omitempty"`

	// 创建队列的用户。
	Owner *string `json:"owner,omitempty"`

	// 创建队列的时间。是单位为“毫秒”的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 队列类型。sql/general/all, 如果不指定，默认为“sql”。
	QueueType *string `json:"queueType,omitempty"`

	// 与队列绑定的最小计算单元个数。设置值当前只支持16，64，256。
	CuCount *int32 `json:"cuCount,omitempty"`

	// 队列的收费模式。 “1”表示按照CU时收费。 “2”表示按照包年包月收费。
	ChargingMode *int32 `json:"chargingMode,omitempty"`

	// 队列的资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 队列类型。 0：共享队列 1：专属队列
	ResourceMode *int32 `json:"resource_mode,omitempty"`

	// 企业项目ID。\"0”表示default，即默认的企业项目。 说明： 开通了企业管理服务的用户可设置该参数绑定指定的项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源类型。 vm：ecf集群 container：容器化集群（k8s）
	ResourceType *string `json:"resource_type,omitempty"`

	// 队列的规格大小。对于包周期队列，表示包周期部分的CU值；对于按需队列，表示用户购买队列时的初始值。
	CuSpec *int32 `json:"cu_spec,omitempty"`

	// 当前队列弹性扩缩容的CU值上限。
	CuScaleOutLimit *int32 `json:"cu_scale_out_limit,omitempty"`

	// 当前队列弹性扩缩容的CU值下限。
	CuScaleInLimit *int32 `json:"cu_scale_in_limit,omitempty"`

	// 弹性资源池名称。
	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`

	// 队列支持的Spark版本。
	SupportSparkVersions *[]string `json:"support_spark_versions,omitempty"`

	// 队列默认的Spark版本。
	DefaultSparkVersion *string `json:"default_spark_version,omitempty"`

	// 队列支持的HetuEngine版本。
	SupportHetuEngineVersions *[]string `json:"support_hetu_engine_versions,omitempty"`

	// 队列默认的HetuEngine版本。
	DefaultHetuEngineVersion *string `json:"default_hetu_engine_version,omitempty"`

	// 队列支持的Flink SQL版本。
	SupportFlinkSqlVersions *[]string `json:"support_flink_sql_versions,omitempty"`

	// 队列默认的Flink SQL版本。
	DefaultFlinkSqlVersion *string `json:"default_flink_sql_version,omitempty"`

	// 队列支持的Flink JAR版本。
	SupportFlinkJarVersions *[]string `json:"support_flink_jar_versions,omitempty"`

	// 队列默认的Flink JAR版本。
	DefaultFlinkJarVersion *string `json:"default_flink_jar_version,omitempty"`
	HttpStatusCode         int     `json:"-"`
}

func (o ShowQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueResponse struct{}"
	}

	return strings.Join([]string{"ShowQueueResponse", string(data)}, " ")
}
