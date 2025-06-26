package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Queue 查询所有队列的响应参数。
type Queue struct {

	// 参数解释: 队列ID 示例: 10 约束限制:  无 取值范围: 无 默认取值: 无
	QueueId *int64 `json:"queue_id,omitempty"`

	// 参数解释: 队列名称 示例: datasource_connection 约束限制:  无 取值范围: 无 默认取值: 无
	QueueName *string `json:"queue_name,omitempty"`

	// 参数解释: 队列描述信息 示例: des 约束限制:  无 取值范围: 无 默认取值: 无
	Description *string `json:"description,omitempty"`

	// 参数解释: 创建队列的用户 示例: ei_dlics_c00228924 约束限制:  无 取值范围: 无 默认取值: 无
	Owner *string `json:"owner,omitempty"`

	// 参数解释: 引擎 示例: spark 约束限制:  无 取值范围: spark, hetuEngine 默认取值: 无
	Engine *QueueEngine `json:"engine,omitempty"`

	// 参数解释: 队列已使用的cu 示例: 6.0 约束限制:  无 取值范围: 大于等于0 默认取值: 无
	UsedCu *float32 `json:"used_cu,omitempty"`

	// 参数解释: 支持的flink版本列表 示例: [1.12, 1.15] 约束限制:  无 取值范围: 无 默认取值: 无
	SupportOpensourceFlinkVersions *[]string `json:"support_opensource_flink_versions,omitempty"`

	// 参数解释: 创建队列的时间。是单位为“毫秒”的时间戳 示例: 1553168198000 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CreateTime *int64 `json:"create_time,omitempty"`

	// 参数解释: 队列的类型 示例: sql 约束限制:  无 取值范围: sql, general, all 默认取值: all
	QueueType *QueueQueueType `json:"queue_type,omitempty"`

	// 参数解释: 与该队列绑定的计算单元数，即当前队列的CU数 示例: 16 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CuCount *int32 `json:"cu_count,omitempty"`

	// 参数解释: 队列的收费模式。 “1”表示按照CU时收费。 “2”表示按照包年包月收费。 示例: 16 约束限制:  无 取值范围: 1, 2 默认取值: 无
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 参数解释: 队列的资源ID 示例: 26afb850-d3c9-42c1-81c0-583d1163e80f 约束限制:  无 取值范围: 无 默认取值: 无
	ResourceId *string `json:"resource_id,omitempty"`

	// 参数解释: 企业项目ID。0”表示default，即默认的企业项目。说明：开通了企业管理服务的用户可设置该参数绑定指定的项目。 示例: 0 约束限制:  无 取值范围: 无 默认取值: 无
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 参数解释: 队列的虚拟私有云（VPC）的网段。建议使用网段：10.0.0.0/8~28，172.16.0.0/12~28，192.168.0.0/16~28 示例: 10.0.0.0/8 约束限制:  符合网段格式的字符串 取值范围: 无 默认取值: 无
	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`

	// 参数解释: 管理子网的网段 示例: 10.23.128.0/24 约束限制:  符合网段格式的字符串 取值范围: 无 默认取值: 无
	CidrInMgntsubnet *string `json:"cidr_in_mgntsubnet,omitempty"`

	// 参数解释: 子网网段 示例: 10.23.128.0/24 约束限制:  符合网段格式的字符串 取值范围: 无 默认取值: 无
	CidrInSubnet *string `json:"cidr_in_subnet,omitempty"`

	// 参数解释: 队列类型。0：共享队列,1：专属队列 示例: 0 约束限制:  无 取值范围: 0,1 默认取值: 无
	ResourceMode *int32 `json:"resource_mode,omitempty"`

	// 参数解释: 队列计算资源的cpu架构 示例: 0 约束限制:  符合cpu架构格式的字符串 取值范围: 无 默认取值: 无
	Platform *string `json:"platform,omitempty"`

	// 参数解释: 是否在重启状态。默认值为“false” 示例: false 约束限制:  无 取值范围: true,false 默认取值: false
	IsRestarting *bool `json:"is_restarting,omitempty"`

	// 参数解释: 队列的标签信息，目前只支持设置跨az配置，multi_az=2 示例: {\\\"multi_az\\\":\\\"2\\\"} 约束限制:  符合Json格式的字符串 取值范围: 无 默认取值: 无
	Labels *string `json:"labels,omitempty"`

	// 参数解释: 队列特性 示例: basic 约束限制:  无 取值范围: basic（基础型） ai（AI增强型，仅SQL的x86_64专属队列支持选择） 默认取值: basic
	Feature *string `json:"feature,omitempty"`

	// 参数解释: 队列所属资源类型 示例: vm 约束限制:  无 取值范围: vm container 默认取值: 无
	QueueResourceType *string `json:"queue_resource_type,omitempty"`

	// 参数解释: 队列的规格大小。对于包周期队列，表示包周期部分的CU值；对于按需队列，表示用户购买队列时的初始值 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CuSpec *int32 `json:"cu_spec,omitempty"`

	// 参数解释: 当前队列弹性扩缩容的CU值上限 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CuScaleOutLimit *int32 `json:"cu_scale_out_limit,omitempty"`

	// 参数解释: 当前队列弹性扩缩容的CU值下限 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CuScaleInLimit *int32 `json:"cu_scale_in_limit,omitempty"`

	// 参数解释: 弹性资源池名称 示例: dli_pool_0509 约束限制:  无 取值范围: 无 默认取值: 无
	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`

	// 参数解释: 队列支持的Spark版本 示例: [2.4.5] 约束限制:  无 取值范围: 无 默认取值: 无
	SupportSparkVersions *[]string `json:"support_spark_versions,omitempty"`

	// 参数解释: 队列默认的Spark版本 示例: 2.4.5 约束限制:  无 取值范围: 无 默认取值: 无
	DefaultSparkVersion *string `json:"default_spark_version,omitempty"`

	// 参数解释: 队列支持的HetuEngine版本 示例: [2.1.0] 约束限制:  无 取值范围: 无 默认取值: 无
	SupportHetuEngineVersions *[]string `json:"support_hetu_engine_versions,omitempty"`

	// 参数解释: 队列默认的HetuEngine版本 示例: 2.1.0 约束限制:  无 取值范围: 无 默认取值: 无
	DefaultHetuEngineVersion *string `json:"default_hetu_engine_version,omitempty"`

	// 参数解释: 队列支持的Flink SQL版本 示例: [1.17] 约束限制:  无 取值范围: 无 默认取值: 无
	SupportFlinkSqlVersions *[]string `json:"support_flink_sql_versions,omitempty"`

	// 参数解释: 队列默认的Flink SQL版本 示例: 1.17 约束限制:  无 取值范围: 无 默认取值: 无
	DefaultFlinkSqlVersion *string `json:"default_flink_sql_version,omitempty"`

	// 参数解释: 队列支持的Flink JAR版本 示例: [1.17] 约束限制:  无 取值范围: 无 默认取值: 无
	SupportFlinkJarVersions *[]string `json:"support_flink_jar_versions,omitempty"`

	// 参数解释: 队列默认的Flink JAR版本 示例: 1.17 约束限制:  无 取值范围: 无 默认取值: 无
	DefaultFlinkJarVersion *string `json:"default_flink_jar_version,omitempty"`
}

func (o Queue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Queue struct{}"
	}

	return strings.Join([]string{"Queue", string(data)}, " ")
}

type QueueEngine struct {
	value string
}

type QueueEngineEnum struct {
	SPARK       QueueEngine
	HETU_ENGINE QueueEngine
}

func GetQueueEngineEnum() QueueEngineEnum {
	return QueueEngineEnum{
		SPARK: QueueEngine{
			value: "spark",
		},
		HETU_ENGINE: QueueEngine{
			value: "hetuEngine",
		},
	}
}

func (c QueueEngine) Value() string {
	return c.value
}

func (c QueueEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueueEngine) UnmarshalJSON(b []byte) error {
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

type QueueQueueType struct {
	value string
}

type QueueQueueTypeEnum struct {
	SQL     QueueQueueType
	GENERAL QueueQueueType
	ALL     QueueQueueType
}

func GetQueueQueueTypeEnum() QueueQueueTypeEnum {
	return QueueQueueTypeEnum{
		SQL: QueueQueueType{
			value: "sql",
		},
		GENERAL: QueueQueueType{
			value: "general",
		},
		ALL: QueueQueueType{
			value: "all",
		},
	}
}

func (c QueueQueueType) Value() string {
	return c.value
}

func (c QueueQueueType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueueQueueType) UnmarshalJSON(b []byte) error {
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
