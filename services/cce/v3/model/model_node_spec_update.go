package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeSpecUpdate struct {

	// 支持给创建出来的节点加Taints来设置反亲和性，taints配置不超过20条。默认值为空。每条Taints包含以下3个参数：  - Key：必须以字母或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀。 - Value：必须以字符或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符。 - Effect：只可选NoSchedule，PreferNoSchedule或NoExecute。  示例：  ``` \"taints\": [{   \"key\": \"status\",   \"value\": \"unavailable\",   \"effect\": \"NoSchedule\" }, {   \"key\": \"looks\",   \"value\": \"bad\",   \"effect\": \"NoSchedule\" }] ``` > 参数未指定或者为空数组时将删除节点池的自定义Taints
	Taints []Taint `json:"taints"`

	// 格式为key/value键值对。键值对个数不超过20条。默认值为空。 - Key：必须以字母或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符；另外可以使用DNS子域作为前缀，例如example.com/my-key，DNS子域最长253个字符。 - Value：可以为空或者非空字符串，非空字符串必须以字符或数字开头，可以包含字母、数字、连字符、下划线和点，最长63个字符。  示例： ``` \"k8sTags\": {   \"key\": \"value\" } ``` > 参数未指定或者为空对象时将删除节点池的自定义K8s标签
	K8sTags map[string]string `json:"k8sTags"`

	// 云服务器标签，键必须唯一，CCE支持的最大用户自定义标签数量依region而定，自定义标签数上限为8个。默认值为空。 > 参数未指定或者为空数组时将删除节点池的自定义云服务器标签
	UserTags []UserTag `json:"userTags"`

	// 自定义初始化标记。  CCE节点在初始化完成之前，会打上初始化未完成污点（node.cloudprovider.kubernetes.io/uninitialized）防止pod调度到节点上。  cce支持自定义初始化标记，在接收到initializedConditions参数后，会将参数值转换成节点标签，随节点下发，例如：cloudprovider.openvessel.io/inject-initialized-conditions=CCEInitial_CustomedInitial。  当节点上设置了此标签，会轮询节点的status.Conditions，查看conditions的type是否存在标记名，如CCEInitial、CustomedInitial标记，如果存在所有传入的标记，且状态为True，认为节点初始化完成，则移除初始化污点。  默认值为空。  - 必须以字母、数字组成，长度范围1-20位。 - 标记数量不超过2个
	InitializedConditions *[]string `json:"initializedConditions,omitempty"`

	Login *Login `json:"login,omitempty"`

	// 服务器企业项目ID。CCE服务不实现EPS相关特性，该字段仅用于同步服务器企业项目ID。 创建节点/节点池场景：可指定已存在企业项目，当取值为空时，该字段继承集群企业项目属性。 更新节点池场景：配置修改后仅会对新增节点的服务器生效，存量节点需前往EPS界面迁移。 如果更新时不指定值，不会更新该字段。 当该字段为空时，返回集群企业项目。
	ServerEnterpriseProjectID *string `json:"serverEnterpriseProjectID,omitempty"`

	NodeNicSpecUpdate *NodeSpecUpdateNodeNicSpecUpdate `json:"nodeNicSpecUpdate,omitempty"`

	// **参数解释**： 指定节点安全加固类型，当前仅支持HCE2.0镜像等保2.0三级安全加固。 等保加固会对身份鉴别、访问控制、安全审计、入侵防范、恶意代码防范进行检查并加固。详情请参见[Huawei Cloud EulerOS 2.0等保2.0三级版镜像概述](https://support.huaweicloud.com/productdesc-hce/hce_sec_0001.html)。 若未指定此参数，则尝试用原有的值补全。如：原先HCE2.0镜像已配置安全加固，更新节点池时未指定此参数，则仍旧保持安全加固配置，若要取消，需显式指定参数值为\"null\"。 **约束限制**： 不涉及 **取值范围**： 取值范围：['null', cybersecurity]; **默认取值**： 不涉及
	SecurityReinforcementType *NodeSpecUpdateSecurityReinforcementType `json:"securityReinforcementType,omitempty"`
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
