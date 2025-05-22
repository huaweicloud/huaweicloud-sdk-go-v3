package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterWithExistedNodesRequestBody **参数解释**： 节点扩容请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ResizeClusterWithExistedNodesRequestBody struct {
	ScaleOut *ScaleOut `json:"scale_out"`

	// **参数解释**： 是否强制备份。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ForceBackup *bool `json:"force_backup,omitempty"`

	// **参数解释**： 扩容模式，不传默认离线read-only。 **约束限制**： 在线模式在大部分低版本集群不支持，请在联系运维人员确认后方才可用。 **取值范围**： read-only：离线模式 insert：在线模式 **默认取值**： 不涉及。
	Mode *string `json:"mode,omitempty"`

	// **参数解释**： 逻辑集群名称。 **约束限制**： 不涉及。 **取值范围**： 非逻辑集群模式下该字段不填，逻辑集群模式下不传默认elastic_group。 **默认取值**： elastic_group
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// **参数解释**： 是否是使用已添加的空闲节点进行扩容。 **约束限制**： 不涉及。 **取值范围**： true：使用空闲节点扩容 false：不使用空闲节点扩容 **默认取值**： false
	ExpandWithExistedNode bool `json:"expand_with_existed_node"`

	// **参数解释**： 扩容完成后是否自动启动重分布，默认true。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： true
	AutoRedistribute *bool `json:"auto_redistribute,omitempty"`

	RedisConf *RedisConfReq `json:"redis_conf,omitempty"`
}

func (o ResizeClusterWithExistedNodesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterWithExistedNodesRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeClusterWithExistedNodesRequestBody", string(data)}, " ")
}
