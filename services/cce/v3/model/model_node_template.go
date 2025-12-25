package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeTemplate struct {

	// **参数解释**： 节点的规格，CCE支持的节点规格请参考[节点规格说明](cce_02_0368.xml)获取。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**： 待创建节点所在的可用区，需要指定可用区（AZ）的名称，通过api创建节点不支持随机可用区。 [CCE支持的可用区请参考[地区和终端节点](https://console.huaweicloud.com/apiexplorer/#/endpoint/CCE)。](tag:hws) [CCE支持的可用区请参考[地区和终端节点](https://console-intl.huaweicloud.com/apiexplorer/#/endpoint/CCE)。](tag:hws_hk) **约束限制**： 创建节点池并设置伸缩组时，该参数不允许填写为random。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Az *string `json:"az,omitempty"`

	// **参数解释**： 节点的操作系统类型。具体支持的操作系统请参见[节点操作系统说明](node-os.xml)。 **约束限制**： - 若在创建节点时未指定该配置，CCE会根据集群版本自动选择支持的OS版本。 - 若当前集群版本不支持该OS类型，则会自动替换为当前集群版本支持的同系列OS类型。 - 若在创建节点时指定了extendParam中的alpha.cce/NodeImageID参数，节点将使用私有镜像，则该参数为非必选参数。 [- 若在创建节点时指定了extendParam中的securityReinforcementType参数为cybersecurity，节点将开启安全等保加固功能，则节点的操作系统类型必须使用HCE2.0。当用户未配置私有镜像时，该参数必须为“Huawei Cloud EulerOS 2.0”；当用户配置了私有镜像且私有镜像操作系统类型为HCE2.0，则该参数为非必选参数。](tag:hws)  **取值范围**： 不涉及 **默认取值**： 不涉及
	Os *string `json:"os,omitempty"`

	Login *Login `json:"login,omitempty"`

	RootVolume *Volume `json:"rootVolume,omitempty"`

	// **参数解释**： 节点的数据盘参数。针对专属云节点，参数解释与rootVolume一致。 **约束限制**： - 磁盘挂载上限为虚拟机不超过16块，裸金属不超过10块。在此基础上还受限于虚拟机/裸金属规格可挂载磁盘数上限。（目前支持通过控制台和API为CCE节点添加多块数据盘）。 - 如果数据盘正供容器运行时和Kubelet组件使用，则不可被卸载，否则将导致节点不可用。 - 仅在选择系统盘作为系统组件存储磁盘时，允许为空。
	DataVolumes *[]Volume `json:"dataVolumes,omitempty"`

	Storage *Storage `json:"storage,omitempty"`

	PublicIP *NodeEipSpec `json:"publicIP,omitempty"`

	NodeNicSpec *NodeNicSpec `json:"nodeNicSpec,omitempty"`

	// **参数解释**： 批量创建时节点的个数。 **约束限制**： - 作用于节点池时该项可以不填写。 - 创建、更新节点池场景返回中无该参数。 - 创建节点时该参数为必填参数  **取值范围**： 必须为大于等于1，小于等于最大限额的正整数。 **默认取值**： 不涉及
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 节点的计费模式。 **约束限制**： 不涉及 **取值范围**： -  0: 按需付费 [- 1: 包周期](tag:hws,hws_hk) [- 2: 已废弃：自动付费包周期](tag:hws,hws_hk)  **默认取值**： 0
	BillingMode *NodeTemplateBillingMode `json:"billingMode,omitempty"`

	// **参数解释**： 支持给创建出来的节点加Taints来设置反亲和性。每条Taints包含以下3个参数：  - Key：必须以字母或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀。 - Value：必须以字符或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符。 - Effect：只可选NoSchedule，PreferNoSchedule或NoExecute。 字段使用场景：在节点创建场景下，支持指定初始值，查询时不返回该字段；在节点池场景下，创建节点池时节点模板参数中支持指定初始值，查询时支持返回该字段；在其余场景下，查询时都不会返回该字段。  示例：  ``` \"taints\": [{   \"key\": \"status\",   \"value\": \"unavailable\",   \"effect\": \"NoSchedule\" }, {   \"key\": \"looks\",   \"value\": \"bad\",   \"effect\": \"NoSchedule\" }] ```  **约束限制**： taints配置不超过20条。
	Taints *[]Taint `json:"taints,omitempty"`

	// **参数解释：** 该参数用于控制创建节点时 **post-install脚本执行完成前允许节点调度** 行为。当该参数未设置或者为false时，在kubernetes节点就绪时，容器即可被调度到可用节点。当该参数为true时，在kubernetes节点就绪时且post-install脚本执行完成时，容器才可被调度到可用节点。  **约束限制：** 不涉及  **取值范围：** - false：在kubernetes节点就绪时，容器即可被调度到可用节点。           - true：在kubernetes节点就绪时且post-install脚本执行完成时，容器才可被调度到可用节点。  **默认取值：** false
	WaitPostInstallFinish *bool `json:"waitPostInstallFinish,omitempty"`

	// **参数解释**： 格式为key/value键值对。 - Key：必须以字母或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀，例如example.com/my-key，DNS子域最长253个字符。 - Value：可以为空或者非空字符串，非空字符串必须以字符或数字开头和结尾，可以包含字母、数字、连字符、下划线和点，最长63个字符。 字段使用场景：在节点创建场景下，支持指定初始值，查询时不返回该字段；在节点池场景下，创建节点池时节点模板参数中支持指定初始值，查询时支持返回该字段；在其余场景下，查询时都不会返回该字段。   示例： ``` \"k8sTags\": {   \"key\": \"value\" } ```  **约束限制**： 键值对个数不超过20条。
	K8sTags map[string]string `json:"k8sTags,omitempty"`

	// **参数解释**： 云服务器组ID，若指定，将节点创建在该云服务器组下。 **约束限制**： 创建、更新节点池时该配置不会生效，若要保持节点池中的节点都在同一个云服务器组内，请在节点池 nodeManagement 字段中配置。 **取值范围**： 不涉及 **默认取值**： 不涉及
	EcsGroupId *string `json:"ecsGroupId,omitempty"`

	// **参数解释**： 云服务器故障域，将节点创建在指定故障域下。 **约束限制**： 必须同时指定故障域策略的云服务器ID，且需要开启故障域特性开关。 **取值范围**： 不涉及 **默认取值**： 不涉及
	FaultDomain *string `json:"faultDomain,omitempty"`

	// **参数解释**： 指定DeH主机的ID，将节点调度到自己的DeH上。 **约束限制**： 创建节点池添加节点时不支持该参数。 **取值范围**： 不涉及 **默认取值**： 不涉及
	DedicatedHostId *string `json:"dedicatedHostId,omitempty"`

	// **参数解释**： 是否CCE Turbo集群节点。 **约束限制**： 创建节点池添加节点时不支持该参数。 **取值范围**： 不涉及 **默认取值**： 不涉及
	OffloadNode *bool `json:"offloadNode,omitempty"`

	// **参数解释**： 节点来源是否为纳管节点。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	IsStatic *bool `json:"isStatic,omitempty"`

	// **参数解释**： 云服务器标签（资源标签）。字段使用场景：在节点创建场景下，支持指定初始值，查询时不返回该字段；在节点池场景下，创建节点池时节点模板参数中支持指定初始值，查询时支持返回该字段；在其余场景下，查询时都不会返回该字段。 **约束限制**： - 键必须唯一，CCE支持的最大用户自定义标签数量依region而定，自定义标签数上限为8个。
	UserTags *[]UserTag `json:"userTags,omitempty"`

	Runtime *Runtime `json:"runtime,omitempty"`

	// **参数解释**： 自定义初始化标记，默认值为空。  CCE节点在初始化完成之前，会打上初始化未完成污点（node.cloudprovider.kubernetes.io/uninitialized）防止pod调度到节点上。用户在创建节点时，可以通过设置initializedConditions参数，控制污点的移除时间（默认不设置超时时间）。  使用示例如下： 1. 创建节点，传入参数 \"initializedConditions\": [\"CCEInitial\", \"CustomedInitial\"]； 2. 用户在执行完自定义初始化操作后，调用k8s接口（例如PATCH /v1/nodes/{node_ip}/status）更新节点的conditions，插入type为CCEInitial、CustomedInitial的两个标记，状态为True，如下所示：   ```   status:     conditions:     - type: CCEInitial       status: 'True'     - type: CustomedInitial       status: 'True'   ``` 3. CCE组件轮询节点的status.Conditions，查看是否存在type为CCEInitial、CustomedInitial的condition，若存在且status字段值为True，认为节点初始化完成，则移除初始化污点； 4. initializedConditions支持设置超时时间，用户可以在创节点时传入，如：\"initializedConditions\": [\"CCEInitial:15m\", \"CustomedInitial:15m\"]，表示超时时间为15分钟，达到超时时间后，当CCE组件轮询到节点时会自动忽略初始化condition，移除初始化污点。  **约束限制**： - initializedConditions中超时时间取值范围为[1-99]秒 - 必须以字母、数字组成，长度范围1-20位。 - 标记数量不超过2个。 - 超时时间仅支持分钟(m)单位。
	InitializedConditions *[]string `json:"initializedConditions,omitempty"`

	ExtendParam *NodeExtendParam `json:"extendParam,omitempty"`

	HostnameConfig *HostnameConfig `json:"hostnameConfig,omitempty"`

	// **参数解释**： 服务器企业项目ID。CCE服务不实现EPS相关特性，该字段仅用于同步服务器企业项目ID。 **约束限制**： 创建节点/节点池场景：可指定已存在企业项目，当取值为空时，该字段继承集群企业项目属性。 更新节点池场景：配置修改后仅会对新增节点的服务器生效，存量节点需前往EPS界面迁移。 **取值范围**： 不涉及 **默认取值**： 如果更新时不指定值，不会更新该字段。 当该字段为空时，返回集群企业项目。
	ServerEnterpriseProjectID *string `json:"serverEnterpriseProjectID,omitempty"`

	// **参数解释**： 表示节点所属分区。分区可以选择中心云[或者[边缘小站](https://support.huaweicloud.com/usermanual-cloudpond/ies_02_0401.html)。](tag:hws)[或者[边缘小站](https://support.huaweicloud.com/intl/zh-cn/usermanual-cloudpond/ies_02_0401.html)。](tag:hws_hk) **约束限制**： 仅开启了对分布式云支持的Turbo集群支持指定该字段。 **取值范围**： - center: 中心云 - 边缘小站的可用区ID  **默认取值**： 不涉及
	Partition *string `json:"partition,omitempty"`

	// **参数解释：** 覆盖节点默认组件配置。  [当前支持的可配置组件及其参数详见[配置管理](https://support.huaweicloud.com/usermanual-cce/cce_10_0652.htmll)。](tag:hws) [当前支持的可配置组件及其参数详见[配置管理](https://support.huaweicloud.com/intl/zh-cn/usermanual-cce/cce_10_0652.html)。](tag:hws_hk) **约束限制：** 若指定了不支持的组件或组件不支持的参数，该配置项将被忽略。
	ConfigurationsOverride *[]PackageConfiguration `json:"configurationsOverride,omitempty"`

	NodeNameTemplate *NodeSpecNodeNameTemplate `json:"nodeNameTemplate,omitempty"`
}

func (o NodeTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeTemplate struct{}"
	}

	return strings.Join([]string{"NodeTemplate", string(data)}, " ")
}

type NodeTemplateBillingMode struct {
	value int32
}

type NodeTemplateBillingModeEnum struct {
	E_0 NodeTemplateBillingMode
	E_1 NodeTemplateBillingMode
}

func GetNodeTemplateBillingModeEnum() NodeTemplateBillingModeEnum {
	return NodeTemplateBillingModeEnum{
		E_0: NodeTemplateBillingMode{
			value: 0,
		}, E_1: NodeTemplateBillingMode{
			value: 1,
		},
	}
}

func (c NodeTemplateBillingMode) Value() int32 {
	return c.value
}

func (c NodeTemplateBillingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeTemplateBillingMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
