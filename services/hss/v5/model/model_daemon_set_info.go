package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DaemonSetInfo daemonset基本信息
type DaemonSetInfo struct {

	// **参数解释**: daemonset名称。 **取值范围**: 字符长度0-256位
	Name *string `json:"name,omitempty"`

	// **参数解释**: 命名空间名称。 **取值范围**: 字符长度0-256位
	NamespaceName *string `json:"namespace_name,omitempty"`

	// **参数解释**: 集群的唯一标识ID，用于唯一确定DaemonSet所属的集群实例。 **取值范围**: 符合标准UUID格式，字符串长度36位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**: DaemonSet所属集群的部署类型，用于区分不同云厂商或部署模式的集群。 **取值范围**: k8s（原生集群）、cce（CCE集群）、ali（阿里云集群）、tencent（腾讯云集群）、azure（微软云集群）、aws（亚马逊集群）、self_built_hw（华为云自建集群）、self_built_idc（IDC自建集群）
	ClusterType *string `json:"cluster_type,omitempty"`

	// **参数解释**: 集群名称。 **取值范围**: 字符串长度0-256
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**: DaemonSet的运行状态，用于标识当前实例的运行健康度； **取值范围**: Running（正常运行）、Failed（存在异常）
	Status *string `json:"status,omitempty"`

	// **参数解释**: 当前DaemonSet对应的Pod实例总数量； **取值范围**: 非负整数，单位为个，最小值0，无上限（取决于集群资源）
	PodsNum *int32 `json:"pods_num,omitempty"`

	// **参数解释**: DaemonSet实例所使用的容器镜像的完整名称，包含镜像仓库、镜像名及标签； **取值范围**: 符合容器镜像命名规范，字符长度1 - 256位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 用于匹配Pod的标签集合，K8s通过该标签关联DaemonSet与对应Pod； **取值范围**: 数组元素数量≥0，每个元素的key为标签名称，val为标签值，key和val字符长度均为1 - 63位
	MatchLabels *[]LabelInfo `json:"match_labels,omitempty"`

	// **参数解释**: DaemonSet资源的创建时间； **取值范围**: 以毫秒级时间戳格式返回，取值为非负长整数，对应UTC时间1970 - 01 - 01起的毫秒数
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o DaemonSetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DaemonSetInfo struct{}"
	}

	return strings.Join([]string{"DaemonSetInfo", string(data)}, " ")
}
