package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReinstallNodeSpec 节点重装配置参数
type ReinstallNodeSpec struct {

	// **参数解释**： 操作系统。指定自定义镜像场景将以IMS镜像的实际操作系统版本为准。请选择当前集群支持的操作系统版本。[例如Huawei Cloud EulerOS 2.0、Ubuntu 22.04、EulerOS 2.9、CentOS 7.6、EulerOS 2.8。](tag:hws,hws_hk) **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Os string `json:"os"`

	Login *Login `json:"login,omitempty"`

	// **参数解释**： 节点名称，重装时指定将修改节点名称，且服务器名称会同步修改。默认以服务器当前名称作为节点名称。 **约束限制**： 命名规则：以小写字母开头，由小写字母、数字、中划线(-)、点(.)组成，长度范围1-56位。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	ServerConfig *ReinstallServerConfig `json:"serverConfig,omitempty"`

	VolumeConfig *ReinstallVolumeConfig `json:"volumeConfig,omitempty"`

	RuntimeConfig *ReinstallRuntimeConfig `json:"runtimeConfig,omitempty"`

	K8sOptions *ReinstallK8sOptionsConfig `json:"k8sOptions,omitempty"`

	Lifecycle *NodeLifecycleConfig `json:"lifecycle,omitempty"`

	// **参数解释**： 自定义初始化标记，默认值为空。  CCE节点在初始化完成之前，会打上初始化未完成污点（node.cloudprovider.kubernetes.io/uninitialized）防止pod调度到节点上。用户在纳管/重置节点时，可以通过设置initializedConditions参数，控制污点的移除时间（默认不设置超时时间）。  使用示例如下： 1. 纳管/重置节点，传入参数 \"initializedConditions\": [\"CCEInitial\", \"CustomedInitial\"]； 2. 用户在执行完自定义初始化操作后，调用k8s接口（例如PATCH /v1/nodes/{node_ip}/status）更新节点的conditions，插入type为CCEInitial、CustomedInitial的两个标记，状态为True； 3. CCE组件轮询节点的status.Conditions，查看是否存在type为CCEInitial、CustomedInitial的condition，若存在且status字段值为True，认为节点初始化完成，则移除初始化污点； 4. initializedConditions支持设置超时时间，用户可以在创节点时传入，如：\"initializedConditions\": [\"CCEInitial:15m\", \"CustomedInitial:15m\"]，表示超时时间为15分钟，达到超时时间后，当CCE组件轮询到节点时会自动忽略初始化condition，移除初始化污点。  节点Conditions示例如下： ``` status:   conditions:   - type: CCEInitial     status: 'True'   - type: CustomedInitial     status: 'True' ```  **约束限制**： - initializedConditions中超时时间取值范围为[1-99]秒 - 必须以字母、数字组成，长度范围1-20位。 - 标记数量不超过2个。 - 超时时间仅支持分钟(m)单位。
	InitializedConditions *[]string `json:"initializedConditions,omitempty"`

	ExtendParam *ReinstallExtendParam `json:"extendParam,omitempty"`

	HostnameConfig *HostnameConfig `json:"hostnameConfig,omitempty"`

	// **参数解释**： 指定节点安全加固类型，当前仅支持HCE2.0镜像等保2.0三级安全加固。 等保加固会对身份鉴别、访问控制、安全审计、入侵防范、恶意代码防范进行检查并加固。[详情请参见[Huawei Cloud EulerOS 2.0等保2.0三级版镜像概述](https://support.huaweicloud.com/productdesc-hce/hce_sec_0001.html)。](tag:hws) 若未指定此参数，则尝试用原有的值补全。如：原先HCE2.0镜像已配置安全加固，更新节点池时未指定此参数，则仍旧保持安全加固配置，若要取消，需显式指定参数值为\"null\"。 **约束限制**： 不涉及 **取值范围**： - 空值：表示不开启等保加固 - cybersecurity：表示开启等保加固  **默认取值**： 不涉及
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
