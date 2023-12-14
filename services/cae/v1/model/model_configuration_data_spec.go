package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationDataSpec 组件配置数据规格详情。
type ConfigurationDataSpec struct {

	// - type为\"rds\"时，配置此参数。 - 参数含义：RDS数据库实例ID。
	RdsId *string `json:"rds_id,omitempty"`

	// - type为\"rds\"时，配置此参数。 - 参数含义：RDS数据库名称。
	RdsDbName *string `json:"rds_db_name,omitempty"`

	// - type为\"rds\"时，配置此参数。 - 参数含义：RDS数据库地址。
	RdsAddress *string `json:"rds_address,omitempty"`

	// - type为\"rds\"时，配置此参数。 - 参数含义：RDS数据库用户名称。
	RdsUsername *string `json:"rds_username,omitempty"`

	// - type为\"rds\"时，配置此参数。 - 参数含义：RDS数据库密码。
	RdsPassword *string `json:"rds_password,omitempty"`

	// - type为\"rds\"时，配置此参数。 - 参数含义：RDS数据库端口
	RdsPort *string `json:"rds_port,omitempty"`

	// - type为\"cse\"时，配置此参数。 - 参数含义：CSE配置中心地址。
	ConfigCenterAddr *string `json:"config_center_addr,omitempty"`

	// - type为\"cse\"时，配置此参数。 - 参数含义：CSE服务注册发现地址。
	ServiceCenterAddr *string `json:"service_center_addr,omitempty"`

	// - type为\"cse\"时，配置此参数。 - 参数含义：CSE引擎ID。
	CseId *string `json:"cse_id,omitempty"`

	// 环境变量配置，常用环境变量如下。 - TZ: 时区设置，东八区可设置为Asia/Shanghai。 - LANG: 语言字符集设置，中文UTF8可设置为zh_CN.UTF-8。
	Envs map[string]string `json:"envs,omitempty"`

	// 弹性公网IP，响应体参数，未配置域名时返回此参数。
	Ip *string `json:"ip,omitempty"`

	// - type为\"access\"时，配置此参数。 - 参数含义：访问方式配置列表。
	Items *[]AccessConfigurationDataItems `json:"items,omitempty"`

	// - type为\"scaling\"时，配置此参数。 - 参数含义：伸缩策略配置最大伸缩个数。
	MaxReplicaCount *int32 `json:"max_replica_count,omitempty"`

	// - type为\"scaling\"时，配置此参数。 - 参数含义：伸缩策略配置触发器列表。
	Triggers *[]ScaleConfigurationDataTrigger `json:"triggers,omitempty"`

	// - type为\"volume\"时，配置此参数。 - 参数含义：云存储配置列表。
	Volumes *[]VolumeConfigurationDataVolume `json:"volumes,omitempty"`

	LivenessProbe *HealthCheckConfigurationLivenessProbe `json:"livenessProbe,omitempty"`

	StartupProbe *HealthCheckConfigurationStartupProbe `json:"startupProbe,omitempty"`

	ReadinessProbe *HealthCheckConfigurationReadinessProbe `json:"readinessProbe,omitempty"`

	PostStart *ConfigurationDataSpecPostStart `json:"postStart,omitempty"`

	PreStop *ConfigurationDataSpecPreStop `json:"preStop,omitempty"`

	// - type为\"log\"时，配置此参数。 - 参数含义：自定义日志路径数组。
	LogPaths *[]string `json:"log_paths,omitempty"`

	// - type为\"apm2\"时，配置此参数。 - 参数含义：探针注入方式。
	Instrumentation *string `json:"instrumentation,omitempty"`

	// - 参数含义：type为\"apm2\"时，性能管理配置应用。
	ApmApplication *string `json:"apm_application,omitempty"`

	// - 参数含义：type为\"apm2\"时，监控系统类别，包括apm2和opentelemetry。
	Type *string `json:"type,omitempty"`

	// - 参数含义：type为\"apm2\"时，apm2组件。
	AppName *string `json:"app_name,omitempty"`

	// - 参数含义：type为\"apm2\"时，apm2实例。
	InstanceName *string `json:"instance_name,omitempty"`

	// - 参数含义：type为\"apm2\"时，apm2环境。
	EnvName *string `json:"env_name,omitempty"`

	// - 参数含义：已废弃，迁移到监控系统，type为\"apm2\"时，性能管理配置升级策略。 - Always，重启自动升级：每次都尝试重新下载镜像。 - IfNotPresent，手动升级: 如果本地有该镜像，则继续使用本地镜像，不下载镜像。
	ImagePullPolicy *string `json:"image_pull_policy,omitempty"`

	// - type为\"apm2\"时，配置此参数。 - 参数含义：已废弃，迁移到监控系统，type为\"apm2\"时，性能管理配置探针版本。
	Version *string `json:"version,omitempty"`

	// - 参数含义：已废弃，迁移到监控系统，type为\"apm2\"时，性能管理配置访问密钥Key。
	AccessKey *string `json:"access_key,omitempty"`

	// - 参数含义：已废弃，迁移到监控系统，type为\"apm2\"时，性能管理配置访问密钥value。
	AccessValue *string `json:"access_value,omitempty"`

	// - 参数含义：已废弃，type为\"apm2\"时，性能管理配置应用，同apm_application。
	Business *string `json:"business,omitempty"`

	// - type为\"customMetric\"时，配置此参数。 - 参数含义：自定义监控指标配置采集路径。
	Path *string `json:"path,omitempty"`

	// - type为\"customMetric\"时，配置此参数。 - 参数含义：自定义监控指标配置采集端口。
	Port *string `json:"port,omitempty"`

	// - type为\"customMetric\"时，配置此参数。 - 参数含义：自定义监控指标配置指标名称。
	Metrics *[]string `json:"metrics,omitempty"`
}

func (o ConfigurationDataSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDataSpec struct{}"
	}

	return strings.Join([]string{"ConfigurationDataSpec", string(data)}, " ")
}
