package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeSpecUpdate struct {

	// **参数解释：** 节点的规格。  **约束限制**： 不涉及 **取值范围：** CCE支持的节点规格请参考[节点规格说明](cce_02_0368.xml)获取。 **默认取值：** 不涉及
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**： 节点的操作系统类型。具体支持的操作系统请参见[节点操作系统说明](node-os.xml)。 **约束限制**： - 若当前集群版本不支持该OS类型，则会自动替换为当前集群版本支持的同系列OS类型。 - 若在创建节点时指定了extendParam中的alpha.cce/NodeImageID参数，可以不填写此参数。 - 该参数缺省时，CCE会根据集群版本自动选择支持的OS版本。 - 创建节点池时，该参数为必选。 - 若创建节点时使用共享磁盘空间，即磁盘初始化配置管理参数使用storage，且StorageGroups中virtualSpaces的name字段指定为share，该参数为必选。  **取值范围**： 不涉及 **默认取值**： 不涉及
	Os *string `json:"os,omitempty"`

	Login *Login `json:"login,omitempty"`

	RootVolumeUpdate *Volume `json:"rootVolumeUpdate,omitempty"`

	// **参数解释**： 节点的数据盘参数。针对专属云节点，参数解释与rootVolume一致。 **约束限制**： 磁盘挂载上限为虚拟机不超过16块，裸金属不超过10块。在此基础上还受限于虚拟机/裸金属规格可挂载磁盘数上限。（目前支持通过控制台和API为CCE节点添加多块数据盘）。 如果数据盘正供容器运行时和Kubelet组件使用，则不可被卸载，否则将导致节点不可用。
	DataVolumesUpdate *[]Volume `json:"dataVolumesUpdate,omitempty"`

	Storage *Storage `json:"storage,omitempty"`

	Runtime *Runtime `json:"runtime,omitempty"`

	// **参数解释**： 支持给创建出来的节点加Taints来设置反亲和性。每条Taints包含以下3个参数：  - Key：必须以字母或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀。 - Value：必须以字符或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符。 - Effect：只可选NoSchedule，PreferNoSchedule或NoExecute。 字段使用场景：在节点创建场景下，支持指定初始值，查询时不返回该字段；在节点池场景下，其中节点模板中支持指定初始值，查询时支持返回该字段；在其余场景下，查询时都不会返回该字段。  示例：  ``` \"taints\": [{   \"key\": \"status\",   \"value\": \"unavailable\",   \"effect\": \"NoSchedule\" }, {   \"key\": \"looks\",   \"value\": \"bad\",   \"effect\": \"NoSchedule\" }] ```  **约束限制**： - taints配置不超过20条。 - 参数未指定或者为空数组时将删除节点池的自定义Taints。 - 更新节点池时，此字段为非必填字段。
	Taints []Taint `json:"taints"`

	// **参数解释**： 格式为key/value键值对。 - Key：必须以字母或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀，例如example.com/my-key，DNS子域最长253个字符。 - Value：可以为空或者非空字符串，非空字符串必须以字符或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符。 字段使用场景：在节点创建场景下，支持指定初始值，查询时不返回该字段；在节点池场景下，其中节点模板中支持指定初始值，查询时支持返回该字段；在其余场景下，查询时都不会返回该字段。   示例： ``` \"k8sTags\": {   \"key\": \"value\" } ```  **约束限制**： - 键值对个数不超过20条。 - 参数未指定或者为空对象时将删除节点池的自定义K8s标签。 - 更新节点池时，此字段为非必填字段。
	K8sTags map[string]string `json:"k8sTags"`

	// **参数解释**： 云服务器组ID，若指定，将节点创建在该云服务器组下。 **约束限制**： 创建节点池时该配置不会生效，若要保持节点池中的节点都在同一个云服务器组内，请在节点池 nodeManagement 字段中配置。 **取值范围**： 不涉及 **默认取值**： 不涉及
	EcsGroupId *string `json:"ecsGroupId,omitempty"`

	// **参数解释**： 云服务器标签（资源标签）。 **约束限制**： - 键必须唯一，CCE支持的最大用户自定义标签数量依region而定，自定义标签数上限为8个。 - 参数未指定或者为空数组时将删除节点池的自定义云服务器标签。 - 更新节点池时，此字段为非必填字段。  **取值范围**： 不涉及 **默认取值**： 不涉及
	UserTags []UserTag `json:"userTags"`

	NodeNameTemplate *NodeSpecUpdateNodeNameTemplate `json:"nodeNameTemplate,omitempty"`

	// **参数解释**： 自定义初始化标记，默认值为空。  CCE节点在初始化完成之前，会打上初始化未完成污点（node.cloudprovider.kubernetes.io/uninitialized）防止pod调度到节点上。用户在创建节点时，可以通过设置initializedConditions参数，控制污点的移除时间（默认不设置超时时间）。  使用示例如下： 1. 创建节点，传入参数 \"initializedConditions\": [\"CCEInitial\", \"CustomedInitial\"]； 2. 更新节点，传入参数 \"initializedConditions\": [\"NodeInitial\"]，节点池新建的节点注册到集群时默认会被设置为不可调度； 3. 用户在执行完自定义初始化操作后，调用k8s接口（例如PATCH /v1/nodes/{node_ip}/status）更新节点的conditions，插入type为CCEInitial、CustomedInitial的两个标记，状态为True，如下所示：   ```   status:     conditions:     - type: CCEInitial       status: 'True'     - type: CustomedInitial       status: 'True'   ``` 4. CCE组件轮询节点的status.Conditions，查看是否存在type为CCEInitial、CustomedInitial的condition，若存在且status字段值为True，认为节点初始化完成，则移除初始化污点； 5. initializedConditions支持设置超时时间，用户可以在创节点时传入，如：\"initializedConditions\": [\"CCEInitial:15m\", \"CustomedInitial:15m\"]，表示超时时间为15分钟，达到超时时间后，当CCE组件轮询到节点时会自动忽略初始化condition，移除初始化污点。  **约束限制**： - initializedConditions中超时时间取值范围为[1-99]秒 - 必须以字母、数字组成，长度范围1-20位。 - 标记数量不超过2个。 - 超时时间仅支持分钟(m)单位。
	InitializedConditions *[]string `json:"initializedConditions,omitempty"`

	// **参数解释**： 服务器企业项目ID。CCE服务不实现EPS相关特性，该字段仅用于同步服务器企业项目ID。 **约束限制**： 创建节点/节点池场景：可指定已存在企业项目，当取值为空时，该字段继承集群企业项目属性。 更新节点池场景：配置修改后仅会对新增节点的服务器生效，存量节点需前往EPS界面迁移。 **取值范围**： 不涉及 **默认取值**： 如果更新时不指定值，不会更新该字段。 当该字段为空时，返回集群企业项目。
	ServerEnterpriseProjectID *string `json:"serverEnterpriseProjectID,omitempty"`

	NodeNicSpecUpdate *NodeSpecUpdateNodeNicSpecUpdate `json:"nodeNicSpecUpdate,omitempty"`

	// **参数解释**： 指定节点安全加固类型，当前仅支持HCE2.0镜像等保2.0三级安全加固。 等保加固会对身份鉴别、访问控制、安全审计、入侵防范、恶意代码防范进行检查并加固。详情请参见[Huawei Cloud EulerOS 2.0等保2.0三级版镜像概述](https://support.huaweicloud.com/productdesc-hce/hce_sec_0001.html)。 若未指定此参数，则尝试用原有的值补全。如：原先HCE2.0镜像已配置安全加固，更新节点池时未指定此参数，则仍旧保持安全加固配置，若要取消，需显式指定参数值为\"null\"。 **约束限制**： 不涉及 **取值范围**： 取值范围：['null', cybersecurity]; **默认取值**： 不涉及
	SecurityReinforcementType *NodeSpecUpdateSecurityReinforcementType `json:"securityReinforcementType,omitempty"`

	ExtendParam *NodePoolUpdateExtendParam `json:"extendParam,omitempty"`
}

func (o NodeSpecUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpecUpdate struct{}"
	}

	return strings.Join([]string{"NodeSpecUpdate", string(data)}, " ")
}

type NodeSpecUpdateSecurityReinforcementType struct {
	value string
}

type NodeSpecUpdateSecurityReinforcementTypeEnum struct {
	NULL          NodeSpecUpdateSecurityReinforcementType
	CYBERSECURITY NodeSpecUpdateSecurityReinforcementType
}

func GetNodeSpecUpdateSecurityReinforcementTypeEnum() NodeSpecUpdateSecurityReinforcementTypeEnum {
	return NodeSpecUpdateSecurityReinforcementTypeEnum{
		NULL: NodeSpecUpdateSecurityReinforcementType{
			value: "null",
		},
		CYBERSECURITY: NodeSpecUpdateSecurityReinforcementType{
			value: "cybersecurity",
		},
	}
}

func (c NodeSpecUpdateSecurityReinforcementType) Value() string {
	return c.value
}

func (c NodeSpecUpdateSecurityReinforcementType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeSpecUpdateSecurityReinforcementType) UnmarshalJSON(b []byte) error {
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
