package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterFlavorRequest Request Object
type ShowClusterFlavorRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 命名空间，一般只填DWS。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Namespace *string `json:"namespace,omitempty"`
}

func (o ShowClusterFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterFlavorRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterFlavorRequest", string(data)}, " ")
}
