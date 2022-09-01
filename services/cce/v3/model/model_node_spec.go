package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeSpec struct {

	// 节点的规格，CCE支持的节点规格请参考[节点规格说明](cce_02_0368.xml)获取。
	Flavor string `json:"flavor" xml:"flavor"`

	// 待创建节点所在的可用区，需要指定可用区（AZ）的名称。 [CCE支持的可用区请参考[地区和终端节点](https://developer.huaweicloud.com/endpoint?CCE)](tag:hws) [CCE支持的可用区请参考[地区和终端节点](https://developer.huaweicloud.com/intl/zh-cn/endpoint?CCE)](tag:hws_hk)
	Az string `json:"az" xml:"az"`

	// 节点的操作系统类型。具体支持的操作系统请参见[节点操作系统说明](node-os.xml)。  > - 系统会根据集群版本自动选择支持的系统版本。当前集群版本不支持该系统类型，则会报错。  > - 若在创建节点时指定了extendParam中的alpha.cce/NodeImageID参数，可以不填写此参数。
	Os *string `json:"os,omitempty" xml:"os"`

	Login *Login `json:"login" xml:"login"`

	RootVolume *Volume `json:"rootVolume" xml:"rootVolume"`

	// 节点的数据盘参数（目前已支持通过控制台为CCE节点添加第二块数据盘）。  针对专属云节点，参数解释与rootVolume一致
	DataVolumes []Volume `json:"dataVolumes" xml:"dataVolumes"`

	Storage *Storage `json:"storage,omitempty" xml:"storage"`

	PublicIP *NodePublicIp `json:"publicIP,omitempty" xml:"publicIP"`

	NodeNicSpec *NodeNicSpec `json:"nodeNicSpec,omitempty" xml:"nodeNicSpec"`

	// 批量创建时节点的个数，必须为大于等于1，小于等于最大限额的正整数。作用于节点池时该项可以不填写。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 节点的计费模式： -  0: 按需付费 [- 1: 包周期](tag:hws,hws_hk) [- 2: 已废弃：自动付费包周期](tag:hws,hws_hk)
	BillingMode *int32 `json:"billingMode,omitempty" xml:"billingMode"`

	// 支持给创建出来的节点加Taints来设置反亲和性，taints配置不超过20条。每条Taints包含以下3个参数：  - Key：必须以字母或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀。 - Value：必须以字符或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符。 - Effect：只可选NoSchedule，PreferNoSchedule或NoExecute。  示例：  ``` \"taints\": [{   \"key\": \"status\",   \"value\": \"unavailable\",   \"effect\": \"NoSchedule\" }, {   \"key\": \"looks\",   \"value\": \"bad\",   \"effect\": \"NoSchedule\" }] ```
	Taints *[]Taint `json:"taints,omitempty" xml:"taints"`

	// 格式为key/value键值对。键值对个数不超过20条。  - Key：必须以字母或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀，例如example.com/my-key， DNS子域最长253个字符。 - Value：可以为空或者非空字符串，非空字符串必须以字符或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符。  示例：  ``` \"k8sTags\": {   \"key\": \"value\" } ```
	K8sTags map[string]string `json:"k8sTags,omitempty" xml:"k8sTags"`

	// 云服务器组ID，若指定，将节点创建在该云服务器组下
	EcsGroupId *string `json:"ecsGroupId,omitempty" xml:"ecsGroupId"`

	// 云服务器故障域，将节点创建在指定故障域下。  >必须同时指定故障域策略的云服务器ID，且需要开启故障域特性开关
	FaultDomain *string `json:"faultDomain,omitempty" xml:"faultDomain"`

	// 指定DeH主机的ID，将节点调度到自己的DeH上。 >创建节点池添加节点时不支持该参数。
	DedicatedHostId *string `json:"dedicatedHostId,omitempty" xml:"dedicatedHostId"`

	// 是否CCE Turbo集群节点 >创建节点池添加节点时不支持该参数。
	OffloadNode *bool `json:"offloadNode,omitempty" xml:"offloadNode"`

	// 节点来源是否为纳管节点
	IsStatic *bool `json:"isStatic,omitempty" xml:"isStatic"`

	// 云服务器标签，键必须唯一，CCE支持的最大用户自定义标签数量依region而定，自定义标签数上限为8个。
	UserTags *[]UserTag `json:"userTags,omitempty" xml:"userTags"`

	Runtime *Runtime `json:"runtime,omitempty" xml:"runtime"`

	ExtendParam *NodeExtendParam `json:"extendParam,omitempty" xml:"extendParam"`
}

func (o NodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpec struct{}"
	}

	return strings.Join([]string{"NodeSpec", string(data)}, " ")
}
