package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveClusterDescriptionInfoRequest Request Object
type SaveClusterDescriptionInfoRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 命名空间。 **约束限制**： 固定值DWS，不填也是DWS。 **取值范围**： DWS **默认取值**： DWS
	Namespace *string `json:"namespace,omitempty"`

	Body *ClusterDescriptionInfo `json:"body,omitempty"`
}

func (o SaveClusterDescriptionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveClusterDescriptionInfoRequest struct{}"
	}

	return strings.Join([]string{"SaveClusterDescriptionInfoRequest", string(data)}, " ")
}
