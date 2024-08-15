package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Queue 查询所有队列的响应参数。
type Queue struct {

	// 队列ID。
	QueueId *int64 `json:"queue_id,omitempty"`

	// 队列名称。
	QueueName *string `json:"queue_name,omitempty"`

	// 队列描述信息。
	Description *string `json:"description,omitempty"`

	// 创建队列的用户。
	Owner *string `json:"owner,omitempty"`

	// 创建队列的时间。是单位为“毫秒”的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 队列的类型。： sql general all 如果不指定，默认为“sql”。
	QueueType *string `json:"queue_type,omitempty"`

	// 队列的实际CU。
	CuCount *int32 `json:"cu_count,omitempty"`

	// 队列的收费模式。 “1”表示按照CU时收费。 “2”表示按照包年包月收费。
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 队列的资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 企业项目ID。0”表示default，即默认的企业项目。 说明： 开通了企业管理服务的用户可设置该参数绑定指定的项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 队列的虚拟私有云（VPC）的网段。建议使用网段：10.0.0.0/8~28，172.16.0.0/12~28，192.168.0.0/16~28。
	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`

	// 管理子网的网段。
	CidrInMgntsubnet *string `json:"cidr_in_mgntsubnet,omitempty"`

	// 子网网段。
	CidrInSubnet *string `json:"cidr_in_subnet,omitempty"`

	// 队列类型。 0：共享队列 1：专属队列
	ResourceMode *int32 `json:"resource_mode,omitempty"`

	// 队列计算资源的cpu架构。
	Platform *string `json:"platform,omitempty"`

	// 是否在重启状态。默认值为“false”。
	IsRestarting *bool `json:"is_restarting,omitempty"`

	// 队列的标签信息，目前只支持设置跨az配置，multi_az=2
	Labels *string `json:"labels,omitempty"`

	// 队列特性。支持以下两种类型：basic：基础型ai：AI增强型（仅SQL的x86_64专属队列支持选择）默认值为“basic”。
	Feature *string `json:"feature,omitempty"`

	// 队列所属资源类型, vm或container。
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
}

func (o Queue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Queue struct{}"
	}

	return strings.Join([]string{"Queue", string(data)}, " ")
}
