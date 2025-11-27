package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Resources 资源
type Resources struct {

	// **参数解释**： 集群id **取值范围**： 不涉及
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群名称 **取值范围**： 不涉及
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 镜像 **取值范围**： 不涉及
	Images *string `json:"images,omitempty"`

	// **参数解释**： 标签 **取值范围**： 不涉及
	Labels *string `json:"labels,omitempty"`

	// **参数解释**： 命名空间 **取值范围**： 不涉及
	Namespace *string `json:"namespace,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
