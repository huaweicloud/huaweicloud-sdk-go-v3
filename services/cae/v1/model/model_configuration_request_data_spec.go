package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfigurationRequestDataSpec 组件配置数据规格详情。
type ConfigurationRequestDataSpec struct {

	// RDS数据库实例ID。  ConfigurationItem.type为\"rds\"时，配置此参数。
	RdsId *string `json:"rds_id,omitempty"`

	// RDS数据库名称。  ConfigurationItem.type为\"rds\"时，配置此参数。
	RdsDbName *string `json:"rds_db_name,omitempty"`

	// RDS数据库地址。  ConfigurationItem.type为\"rds\"时，配置此参数。
	RdsAddress *string `json:"rds_address,omitempty"`

	// RDS数据库用户名称。  ConfigurationItem.type为\"rds\"时，配置此参数。
	RdsUsername *string `json:"rds_username,omitempty"`

	// RDS数据库密码。  ConfigurationItem.type为\"rds\"时，配置此参数。
	RdsPassword *string `json:"rds_password,omitempty"`

	// RDS数据库端口。  ConfigurationItem.type为\"rds\"时，配置此参数。
	RdsPort *string `json:"rds_port,omitempty"`

	// CSE配置中心地址。  ConfigurationItem.type为\"cse\"时，配置此参数。
	ConfigCenterAddr *string `json:"config_center_addr,omitempty"`

	// CSE服务注册发现地址。  ConfigurationItem.type为\"cse\"时，配置此参数。
	ServiceCenterAddr *string `json:"service_center_addr,omitempty"`

	// CSE引擎ID。  ConfigurationItem.type为\"cse\"时，配置此参数。
	CseId *string `json:"cse_id,omitempty"`

	// 环境变量配置。 常用环境变量如下： - TZ: 时区设置，东八区可设置为Asia/Shanghai。 - LANG: 语言字符集设置，中文UTF8可设置为zh_CN.UTF-8。
	Envs map[string]string `json:"envs,omitempty"`

	// 弹性公网IP，响应体参数，未配置域名时返回此参数。
	Ip *string `json:"ip,omitempty"`

	// 访问方式配置列表。  ConfigurationItem.type为\"access\"时，配置此参数。
	Items *[]AccessConfigurationDataItems `json:"items,omitempty"`

	// 伸缩策略配置策略类型。  ConfigurationItem.type为\"scaling\"时，配置此参数。
	ScaleStrategy *ConfigurationRequestDataSpecScaleStrategy `json:"scale_strategy,omitempty"`

	// 伸缩策略配置最大伸缩个数。  ConfigurationItem.type为\"scaling\"时，配置此参数。
	MaxReplicaCount *int32 `json:"max_replica_count,omitempty"`

	// 伸缩策略配置最小伸缩个数。  ConfigurationItem.type为\"scaling\"时，配置此参数。
	MinReplicaCount *int32 `json:"min_replica_count,omitempty"`

	Advanced *ScaleConfigurationDataAdvanced `json:"advanced,omitempty"`

	// 伸缩策略配置触发器列表。  ConfigurationItem.type为\"scaling\"时，配置此参数。
	Triggers *[]ScaleConfigurationDataTrigger `json:"triggers,omitempty"`

	// 云存储配置列表。  ConfigurationItem.type为\"volume\"时，配置此参数。
	Volumes *[]VolumeConfigurationDataVolume `json:"volumes,omitempty"`

	LivenessProbe *HealthCheckConfigurationLivenessProbe `json:"livenessProbe,omitempty"`

	StartupProbe *HealthCheckConfigurationStartupProbe `json:"startupProbe,omitempty"`

	ReadinessProbe *HealthCheckConfigurationReadinessProbe `json:"readinessProbe,omitempty"`

	PostStart *ConfigurationRequestDataSpecPostStart `json:"postStart,omitempty"`

	PreStop *ConfigurationRequestDataSpecPreStop `json:"preStop,omitempty"`

	// 自定义本地磁盘日志路径数组。  ConfigurationItem.type为\"log\"时，配置此参数。
	LogPaths *[]string `json:"log_paths,omitempty"`

	// 自定义云存储日志路径数组。  ConfigurationItem.type为\"log\"时，配置此参数。
	CloudStorageLogPaths *[]CloudStorageLogPathInfo `json:"cloud_storage_log_paths,omitempty"`

	// 探针注入方式。  ConfigurationItem.type为\"apm2\"时，配置此参数。
	Instrumentation *string `json:"instrumentation,omitempty"`

	// 自定义监控指标配置采集路径。  ConfigurationItem.type为\"customMetric\"时，配置此参数。
	Path *string `json:"path,omitempty"`

	// 自定义监控指标配置采集端口。  ConfigurationItem.type为\"customMetric\"时，配置此参数。
	Port *int32 `json:"port,omitempty"`

	// 自定义监控指标配置指标名称。  ConfigurationItem.type为\"customMetric\"时，配置此参数。
	Metrics *[]string `json:"metrics,omitempty"`
}

func (o ConfigurationRequestDataSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationRequestDataSpec struct{}"
	}

	return strings.Join([]string{"ConfigurationRequestDataSpec", string(data)}, " ")
}

type ConfigurationRequestDataSpecScaleStrategy struct {
	value string
}

type ConfigurationRequestDataSpecScaleStrategyEnum struct {
	METRIC ConfigurationRequestDataSpecScaleStrategy
	TIME   ConfigurationRequestDataSpecScaleStrategy
	MIX    ConfigurationRequestDataSpecScaleStrategy
}

func GetConfigurationRequestDataSpecScaleStrategyEnum() ConfigurationRequestDataSpecScaleStrategyEnum {
	return ConfigurationRequestDataSpecScaleStrategyEnum{
		METRIC: ConfigurationRequestDataSpecScaleStrategy{
			value: "metric",
		},
		TIME: ConfigurationRequestDataSpecScaleStrategy{
			value: "time",
		},
		MIX: ConfigurationRequestDataSpecScaleStrategy{
			value: "mix",
		},
	}
}

func (c ConfigurationRequestDataSpecScaleStrategy) Value() string {
	return c.value
}

func (c ConfigurationRequestDataSpecScaleStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationRequestDataSpecScaleStrategy) UnmarshalJSON(b []byte) error {
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
