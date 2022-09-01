package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点重装配置参数
type ReinstallNodeSpec struct {

	// 操作系统。指定自定义镜像场景将以IMS镜像的实际操作系统版本为准。请选择当前集群支持的操作系统版本，例如EulerOS 2.5、CentOS 7.6、EulerOS 2.8。
	Os string `json:"os" xml:"os"`

	Login *Login `json:"login" xml:"login"`

	// 节点名称  > 重装时指定将修改节点名称，且服务器名称会同步修改。默认以服务器当前名称作为节点名称。 > > 命名规则：以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围1-56位，且不能以中划线(-)结尾。
	Name *string `json:"name,omitempty" xml:"name"`

	ServerConfig *ReinstallServerConfig `json:"serverConfig,omitempty" xml:"serverConfig"`

	VolumeConfig *ReinstallVolumeConfig `json:"volumeConfig,omitempty" xml:"volumeConfig"`

	RuntimeConfig *ReinstallRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig"`

	K8sOptions *ReinstallK8sOptionsConfig `json:"k8sOptions,omitempty" xml:"k8sOptions"`

	Lifecycle *NodeLifecycleConfig `json:"lifecycle,omitempty" xml:"lifecycle"`

	ExtendParam *ReinstallExtendParam `json:"extendParam,omitempty" xml:"extendParam"`
}

func (o ReinstallNodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallNodeSpec struct{}"
	}

	return strings.Join([]string{"ReinstallNodeSpec", string(data)}, " ")
}
