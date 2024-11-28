package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReinstallNodeSpec 节点重装配置参数
type ReinstallNodeSpec struct {

	// 操作系统。指定自定义镜像场景将以IMS镜像的实际操作系统版本为准。请选择当前集群支持的操作系统版本，例如EulerOS 2.5、CentOS 7.6、EulerOS 2.8。
	Os string `json:"os"`

	Login *Login `json:"login"`

	// 节点名称 > 重装时指定将修改节点名称，且服务器名称会同步修改。默认以服务器当前名称作为节点名称。 > 命名规则：以小写字母开头，由小写字母、数字、中划线(-)组成，长度范围1-56位。
	Name *string `json:"name,omitempty"`

	ServerConfig *ReinstallServerConfig `json:"serverConfig,omitempty"`

	VolumeConfig *ReinstallVolumeConfig `json:"volumeConfig,omitempty"`

	RuntimeConfig *ReinstallRuntimeConfig `json:"runtimeConfig,omitempty"`

	K8sOptions *ReinstallK8sOptionsConfig `json:"k8sOptions,omitempty"`

	Lifecycle *NodeLifecycleConfig `json:"lifecycle,omitempty"`

	// 自定义初始化标记。  CCE节点在初始化完成之前，会打上初始化未完成污点（node.cloudprovider.kubernetes.io/uninitialized）防止pod调度到节点上。  cce支持自定义初始化标记，在接收到initializedConditions参数后，会将参数值转换成节点标签，随节点下发，例如：cloudprovider.openvessel.io/inject-initialized-conditions=CCEInitial_CustomedInitial。  当节点上设置了此标签，会轮询节点的status.Conditions，查看conditions的type是否存在标记名，如CCEInitial、CustomedInitial标记，如果存在所有传入的标记，且状态为True，认为节点初始化完成，则移除初始化污点。  - 必须以字母、数字组成，长度范围1-20位。 - 标记数量不超过2个
	InitializedConditions *[]string `json:"initializedConditions,omitempty"`

	ExtendParam *ReinstallExtendParam `json:"extendParam,omitempty"`

	HostnameConfig *HostnameConfig `json:"hostnameConfig,omitempty"`

	// **参数解释**： 指定节点安全加固类型，当前仅支持HCE2.0镜像等保2.0三级安全加固。 等保加固会对身份鉴别、访问控制、安全审计、入侵防范、恶意代码防范进行检查并加固。详情请参见[Huawei Cloud EulerOS 2.0等保2.0三级版镜像概述](https://support.huaweicloud.com/productdesc-hce/hce_sec_0001.html)。 若未指定此参数，则尝试用原有的值补全。如：原先HCE2.0镜像已配置安全加固，更新节点池时未指定此参数，则仍旧保持安全加固配置，若要取消，需显式指定参数值为\"null\"。 **约束限制**： 不涉及 **取值范围**： 取值范围：['null', cybersecurity]; **默认取值**： 不涉及
	SecurityReinforcementType *ReinstallNodeSpecSecurityReinforcementType `json:"securityReinforcementType,omitempty"`
}

func (o ReinstallNodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallNodeSpec struct{}"
	}

	return strings.Join([]string{"ReinstallNodeSpec", string(data)}, " ")
}

type ReinstallNodeSpecSecurityReinforcementType struct {
	value string
}

type ReinstallNodeSpecSecurityReinforcementTypeEnum struct {
	NULL          ReinstallNodeSpecSecurityReinforcementType
	CYBERSECURITY ReinstallNodeSpecSecurityReinforcementType
}

func GetReinstallNodeSpecSecurityReinforcementTypeEnum() ReinstallNodeSpecSecurityReinforcementTypeEnum {
	return ReinstallNodeSpecSecurityReinforcementTypeEnum{
		NULL: ReinstallNodeSpecSecurityReinforcementType{
			value: "null",
		},
		CYBERSECURITY: ReinstallNodeSpecSecurityReinforcementType{
			value: "cybersecurity",
		},
	}
}

func (c ReinstallNodeSpecSecurityReinforcementType) Value() string {
	return c.value
}

func (c ReinstallNodeSpecSecurityReinforcementType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReinstallNodeSpecSecurityReinforcementType) UnmarshalJSON(b []byte) error {
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
