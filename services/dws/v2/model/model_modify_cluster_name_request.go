package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyClusterNameRequest Request Object
type ModifyClusterNameRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	Body *ClusterNameReq `json:"body,omitempty"`
}

func (o ModifyClusterNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyClusterNameRequest struct{}"
	}

	return strings.Join([]string{"ModifyClusterNameRequest", string(data)}, " ")
}
