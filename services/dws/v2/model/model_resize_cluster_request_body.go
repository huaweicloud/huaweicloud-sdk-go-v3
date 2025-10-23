package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeClusterRequestBody **参数解释**： 扩容/添加空闲节点操作请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type ResizeClusterRequestBody struct {
	ScaleOut *ScaleOut `json:"scale_out"`

	// **参数解释**： 当前是否仅添加空闲节点。 **约束限制**： 不涉及。 **取值范围**： true：仅添加节点，如需扩容则需要单独操作 false：添加节点并扩容集群 **默认取值**： false
	CreateNodeOnly *bool `json:"create_node_only,omitempty"`

	// **参数解释**： 自动查杀作业等待时间。 **约束限制**： guestAgent插件版本8.2.1及以上才支持。 **取值范围**： 30~1200 **默认取值**： 0，即不限制。
	WaitingForKilling *int32 `json:"waiting_for_killing,omitempty"`

	// **参数解释**： 扩容完成后是否自动启动重分布，默认是。如果设置为false，扩容后不进行重分布，此时集群任务信息处于“待重分布”状态，无法进行其他操作。 **约束限制**： 不涉及。 **取值范围**： true：扩容后立即重分布。 false：扩容后不进行重分布，此时集群任务信息处于“待重分布”状态。 **默认取值**： true
	AutoRedistribute *bool `json:"auto_redistribute,omitempty"`
}

func (o ResizeClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeClusterRequestBody", string(data)}, " ")
}
