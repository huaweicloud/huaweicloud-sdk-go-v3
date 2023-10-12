package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessConfigDeatilUpdate struct {

	// 1.路径必须以/或者字母:\\开头 2.不能包含特殊字符<> ' | \" 且不能只输入/ 3.第一级目录不支持通配符*：不能以/_**   /_*开头 4.**只能出现一次 5.最大数量为10
	Paths *[]string `json:"paths,omitempty"`

	// 1.路径必须以/或者字母:\\开头 2.不能包含特殊字符<> ' | \" 且不能只输入/ 3.第一级目录不支持通配符*：不能以/_**   /_*开头 4.**只能出现一次 5.最大数量为10
	BlackPaths *[]string `json:"black_paths,omitempty"`

	Format *AccessConfigFormatUpdate `json:"format,omitempty"`

	WindowsLogInfo *AccessConfigWindowsLogInfoUpdate `json:"windows_log_info,omitempty"`

	// 标准输出开关，仅CCE接入类型时使用
	Stdout *bool `json:"stdout,omitempty"`

	// 标准输出开关标准错误开关，仅CCE接入类型时使用
	Stderr *bool `json:"stderr,omitempty"`

	// CCE接入类型，仅CCE接入类型时使用
	PathType *AccessConfigDeatilUpdatePathType `json:"pathType,omitempty"`

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
}

func (o AccessConfigDeatilUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigDeatilUpdate struct{}"
	}

	return strings.Join([]string{"AccessConfigDeatilUpdate", string(data)}, " ")
}

type AccessConfigDeatilUpdatePathType struct {
	value string
}

type AccessConfigDeatilUpdatePathTypeEnum struct {
	HOST_FILE        AccessConfigDeatilUpdatePathType
	CONTAINER_STDOUT AccessConfigDeatilUpdatePathType
	CONTAINER_FILE   AccessConfigDeatilUpdatePathType
}

func GetAccessConfigDeatilUpdatePathTypeEnum() AccessConfigDeatilUpdatePathTypeEnum {
	return AccessConfigDeatilUpdatePathTypeEnum{
		HOST_FILE: AccessConfigDeatilUpdatePathType{
			value: "HOST_FILE",
		},
		CONTAINER_STDOUT: AccessConfigDeatilUpdatePathType{
			value: "CONTAINER_STDOUT",
		},
		CONTAINER_FILE: AccessConfigDeatilUpdatePathType{
			value: "CONTAINER_FILE",
		},
	}
}

func (c AccessConfigDeatilUpdatePathType) Value() string {
	return c.value
}

func (c AccessConfigDeatilUpdatePathType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigDeatilUpdatePathType) UnmarshalJSON(b []byte) error {
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
