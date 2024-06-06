package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeGroupV2 struct {

	// 节点组名称，最大长度64，支持大小写英文、数字以及“_”。节点组配置原则如下： - master_node_default_group：Master节点组，所有集群类型均需包含该节点组。 - core_node_analysis_group：分析Core节点组，分析集群、混合集群均需包含该节点组。 - core_node_streaming_group：流式Core节点组，流式集群和混合集群均需包含该节点组。 - task_node_analysis_group：分析Task节点组，分析集群和混合集群可根据需要选择该节点组。 - task_node_streaming_group：流式Task节点组，流式集群、混合集群可根据需要选择该节点组。 - node_group{x}：自定义集群节点组，可根据需要添加多个，最多支持添加9个该节点组。
	GroupName string `json:"group_name"`

	// 节点数量，取值范围0～500，Core与Task节点总数最大为500个。
	NodeNum int32 `json:"node_num"`

	// 节点的实例规格。 例如：c3.4xlarge.2.linux.bigdata。实例规格详细说明请参见[MRS所使用的弹性云服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9006.html)和[MRS所使用的裸金属服务器规格](https://support.huaweicloud.com/api-mrs/mrs_01_9001.html)。 该参数建议从MRS控制台的集群创建页面获取对应区域对应版本所支持的规格。
	NodeSize string `json:"node_size"`

	RootVolume *Volume `json:"root_volume,omitempty"`

	DataVolume *Volume `json:"data_volume,omitempty"`

	// 节点数据磁盘存储数目，取值范围：0～20。
	DataVolumeCount *int32 `json:"data_volume_count,omitempty"`

	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`

	AutoScalingPolicy *AutoScalingPolicy `json:"auto_scaling_policy,omitempty"`

	// 当集群类型为CUSTOM时，该参数必选。可以指定节点组中部署的角色，该参数是一个字符串数组，每个字符串表示一个角色表达式。 角色表达式定义： - 当该角色在节点组所有节点部署时： {role name}，如“DataNode”。 - 当该角色在节点组指定下标节点部署时：{role name}:{index1},{index2}…,{indexN}，如“NameNode:1,2”，下标从1开始计数。 - 部分角色支持多实例部署（即在一个节点部署多个同角色的实例）：{role name}[{instance count}]，如“EsNode[9]” 可选的角色请参考[MRS支持的角色与组件对应表](https://support.huaweicloud.com/api-mrs/mrs_02_0106.html)。
	AssignedRoles *[]string `json:"assigned_roles,omitempty"`
}

func (o NodeGroupV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeGroupV2 struct{}"
	}

	return strings.Join([]string{"NodeGroupV2", string(data)}, " ")
}
