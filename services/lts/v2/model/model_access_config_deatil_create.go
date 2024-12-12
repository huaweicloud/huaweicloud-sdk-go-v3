package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigDeatilCreate 日志接入详细信息
type AccessConfigDeatilCreate struct {

	// 采集路径。 1. 路径必须以/或者字母:\\\\开头 2. 不能包含特殊字符<> ' | \" 且不能只输入/ 3. 第一级目录不支持通配符*：不能以/_**   /_*开头 4.**只能出现一次`` CCE类型中 容器路径和主机路径必填，标准输出不用 最小长度：1 最大长度：128
	Paths *[]string `json:"paths,omitempty"`

	// 采集路径黑名单。 1. 路径必须以/或者字母:\\\\开头 2. 不能包含特殊字符<> ' | \" 且不能只输入/ 3. 第一级目录不支持通配符*：不能以/_**   /_*开头 4.**只能出现一次 最小长度：1 最大长度：128
	BlackPaths *[]string `json:"black_paths,omitempty"`

	Format *AccessConfigFormatCreate `json:"format"`

	WindowsLogInfo *AccessConfigWindowsLogInfoCreate `json:"windows_log_info,omitempty"`

	// 标准输出开关，仅CCE接入类型时使用
	Stdout *bool `json:"stdout,omitempty"`

	// 标准输出开关标准错误开关，仅CCE接入类型时使用
	Stderr *bool `json:"stderr,omitempty"`

	// CCE接入类型，仅CCE接入类型时使用
	PathType *AccessConfigDeatilCreatePathType `json:"pathType,omitempty"`

	// K8s Namespace正则匹配，仅CCE接入类型时使用
	NamespaceRegex *string `json:"namespaceRegex,omitempty"`

	// K8s Pod正则匹配，仅CCE接入类型时使用
	PodNameRegex *string `json:"podNameRegex,omitempty"`

	// K8s 容器名称正则匹配，仅CCE接入类型时使用
	ContainerNameRegex *string `json:"containerNameRegex,omitempty"`

	// 容器 Label白名单，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	IncludeLabels map[string]string `json:"includeLabels,omitempty"`

	// 容器 Label黑名单，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	ExcludeLabels map[string]string `json:"excludeLabels,omitempty"`

	// 环境变量白名单，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	IncludeEnvs map[string]string `json:"includeEnvs,omitempty"`

	// 环境变量黑名单，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	ExcludeEnvs map[string]string `json:"excludeEnvs,omitempty"`

	// 容器 Label日志标签，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	LogLabels map[string]string `json:"logLabels,omitempty"`

	// 环境变量日志标签，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	LogEnvs map[string]string `json:"logEnvs,omitempty"`

	// K8s Label白名单，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	IncludeK8sLabels map[string]string `json:"includeK8sLabels,omitempty"`

	// K8s Label黑名单，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	ExcludeK8sLabels map[string]string `json:"excludeK8sLabels,omitempty"`

	// K8s Label日志标签，最多支持创建30个，keyname不支持重名，仅CCE接入类型时使用
	LogK8s map[string]string `json:"logK8s,omitempty"`

	// 是否允许此文件重复采集
	RepeatCollect *bool `json:"repeat_collect,omitempty"`

	// 系统内置字段：配置日志接入规则时，可以配置系统内置字段，上报日志后，每条日志数据的标签数据中将会有系统字段 采集场景为主机文件的内置字段为：hostName、hostId、hostIP、pathFile、hostIPv6、category、collectTime、__host_group__ 采集场景为K8S集群容器文件的内置字段为：hostName、hostId、hostIP、pathFile、hostIPv6、clusterId、podName、appName、containerName、nameSpace、category、collectTime、__host_group__、serviceID、podIp、clusterName、workloadType 若修改时传入此字段，将覆盖原有配置
	SystemFields *[]string `json:"system_fields,omitempty"`

	// 自定义键值对：配置日志接入规则时，可以配置自定义键值对规则，上报日志后，每条日志数据的标签数据中将会有用户自定义的键值对字段，键值对数量不超过20 键的长度限制为128，允许的字符有a-zA-Z0-9_- 值的长度限制为1024
	CustomKeyValue map[string]string `json:"custom_key_value,omitempty"`

	// 容器 Label白名单，可选为AND，OR，不配置时默认为OR；当存在多个值时的处理逻辑，AND表示同时满足才会生效，OR表示有一项满足就会生效
	IncludeLabelsLogical *string `json:"includeLabelsLogical,omitempty"`

	// 容器 Label黑名单，可选为AND，OR，不配置时默认为OR；当存在多个值时的处理逻辑，AND表示同时满足才会生效，OR表示有一项满足就会生效
	ExcludeLabelsLogical *string `json:"excludeLabelsLogical,omitempty"`

	// K8S Label白名单，可选为AND，OR，不配置时默认为OR；当存在多个值时的处理逻辑，AND表示同时满足才会生效，OR表示有一项满足就会生效
	IncludeK8sLabelsLogical *string `json:"includeK8sLabelsLogical,omitempty"`

	// K8S Label黑名单，可选为AND，OR，不配置时默认为OR；当存在多个值时的处理逻辑，AND表示同时满足才会生效，OR表示有一项满足就会生效
	ExcludeK8sLabelsLogical *string `json:"excludeK8sLabelsLogical,omitempty"`

	// 环境变量白名单，可选为AND，OR，不配置时默认为OR；当存在多个值时的处理逻辑，AND表示同时满足才会生效，OR表示有一项满足就会生效
	IncludeEnvsLogical *string `json:"includeEnvsLogical,omitempty"`

	// 环境变量黑名单，可选为AND，OR，不配置时默认为OR；当存在多个值时的处理逻辑，AND表示同时满足才会生效，OR表示有一项满足就会生效
	ExcludeEnvsLogical *string `json:"excludeEnvsLogical,omitempty"`
}

func (o AccessConfigDeatilCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigDeatilCreate struct{}"
	}

	return strings.Join([]string{"AccessConfigDeatilCreate", string(data)}, " ")
}

type AccessConfigDeatilCreatePathType struct {
	value string
}

type AccessConfigDeatilCreatePathTypeEnum struct {
	HOST_FILE        AccessConfigDeatilCreatePathType
	CONTAINER_STDOUT AccessConfigDeatilCreatePathType
	CONTAINER_FILE   AccessConfigDeatilCreatePathType
}

func GetAccessConfigDeatilCreatePathTypeEnum() AccessConfigDeatilCreatePathTypeEnum {
	return AccessConfigDeatilCreatePathTypeEnum{
		HOST_FILE: AccessConfigDeatilCreatePathType{
			value: "HOST_FILE",
		},
		CONTAINER_STDOUT: AccessConfigDeatilCreatePathType{
			value: "CONTAINER_STDOUT",
		},
		CONTAINER_FILE: AccessConfigDeatilCreatePathType{
			value: "CONTAINER_FILE",
		},
	}
}

func (c AccessConfigDeatilCreatePathType) Value() string {
	return c.value
}

func (c AccessConfigDeatilCreatePathType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigDeatilCreatePathType) UnmarshalJSON(b []byte) error {
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
