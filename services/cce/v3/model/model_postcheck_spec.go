package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostcheckSpec struct {

	// **参数解释：** 集群ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterID *string `json:"clusterID,omitempty"`

	// **参数解释：** 升级前的集群版本 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ClusterVersion *string `json:"clusterVersion,omitempty"`

	// **参数解释：** 当前集群版本 **约束限制：** 不涉及 **取值范围：** CCE支持的集群版本 **默认取值：** 不涉及
	TargetVersion *string `json:"targetVersion,omitempty"`
}

func (o PostcheckSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostcheckSpec struct{}"
	}

	return strings.Join([]string{"PostcheckSpec", string(data)}, " ")
}
