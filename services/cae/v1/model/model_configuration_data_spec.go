package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfigurationDataSpec 组件配置数据规格详情。
type ConfigurationDataSpec struct {

	// 弹性公网IP。
	Ip *string `json:"ip,omitempty"`

	// RDS数据库实例ID。
	RdsId *string `json:"rds_id,omitempty"`

	// RDS数据库名称。
	RdsDbName *string `json:"rds_db_name,omitempty"`

	// RDS数据库地址。
	RdsAddress *string `json:"rds_address,omitempty"`

	// RDS数据库用户名称。
	RdsUsername *string `json:"rds_username,omitempty"`

	// RDS数据库密码。
	RdsPassword *string `json:"rds_password,omitempty"`

	// RDS数据库端口
	RdsPort *string `json:"rds_port,omitempty"`

	// CSE配置中心地址。
	ConfigCenterAddr *string `json:"config_center_addr,omitempty"`

	// CSE服务注册发现地址。
	ServiceCenterAddr *string `json:"service_center_addr,omitempty"`

	// CSE引擎ID。
	CseId *string `json:"cse_id,omitempty"`

	// 环境变量配置，常用环境变量如下。 - TZ: 时区设置，东八区可设置为Asia/Shanghai。 - LANG: 语言字符集设置，中文UTF8可设置为zh_CN.UTF-8。
	Envs map[string]string `json:"envs,omitempty"`

	// 访问方式配置列表。
	Items *[]AccessConfigurationDataItems `json:"items,omitempty"`

	// 伸缩策略配置最大伸缩个数。
	MaxReplicaCount *int32 `json:"max_replica_count,omitempty"`

	// 伸缩策略配置触发器列表。
	Triggers *[]ScaleConfigurationDataTrigger `json:"triggers,omitempty"`

	// 云存储配置列表。
	Volumes *[]VolumeConfigurationDataVolume `json:"volumes,omitempty"`

	LivenessProbe *HealthCheckConfigurationLivenessProbe `json:"livenessProbe,omitempty"`

	StartupProbe *HealthCheckConfigurationStartupProbe `json:"startupProbe,omitempty"`

	ReadinessProbe *HealthCheckConfigurationReadinessProbe `json:"readinessProbe,omitempty"`

	PostStart *ConfigurationDataSpecPostStart `json:"postStart,omitempty"`

	PreStop *ConfigurationDataSpecPreStop `json:"preStop,omitempty"`

	// 自定义日志路径。
	LogPaths *[]interface{} `json:"log_paths,omitempty"`

	// 性能管理配置访问密钥Key。
	AccessKey *string `json:"access_key,omitempty"`

	// 性能管理配置访问密钥value。
	AccessValue *string `json:"access_value,omitempty"`

	// 性能管理配置应用。
	Business *string `json:"business,omitempty"`

	// 性能管理配置升级策略。 - 重启自动升级：每次都尝试重新下载镜像。 - 手动升级: 如果本地有该镜像，则继续使用本地镜像，不下载镜像。
	ImagePullPolicy *ConfigurationDataSpecImagePullPolicy `json:"image_pull_policy,omitempty"`

	// 性能管理配置探针版本。
	Version *string `json:"version,omitempty"`

	// 自定义监控指标配置采集路径。
	Path *string `json:"path,omitempty"`

	// 自定义监控指标配置采集端口。
	Port *string `json:"port,omitempty"`

	// 自定义监控指标配置指标名称。
	Metrics *[]interface{} `json:"metrics,omitempty"`
}

func (o ConfigurationDataSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDataSpec struct{}"
	}

	return strings.Join([]string{"ConfigurationDataSpec", string(data)}, " ")
}

type ConfigurationDataSpecImagePullPolicy struct {
	value string
}

type ConfigurationDataSpecImagePullPolicyEnum struct {
	ALWAYS         ConfigurationDataSpecImagePullPolicy
	IF_NOT_PRESENT ConfigurationDataSpecImagePullPolicy
}

func GetConfigurationDataSpecImagePullPolicyEnum() ConfigurationDataSpecImagePullPolicyEnum {
	return ConfigurationDataSpecImagePullPolicyEnum{
		ALWAYS: ConfigurationDataSpecImagePullPolicy{
			value: "Always",
		},
		IF_NOT_PRESENT: ConfigurationDataSpecImagePullPolicy{
			value: "IfNotPresent",
		},
	}
}

func (c ConfigurationDataSpecImagePullPolicy) Value() string {
	return c.value
}

func (c ConfigurationDataSpecImagePullPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationDataSpecImagePullPolicy) UnmarshalJSON(b []byte) error {
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
