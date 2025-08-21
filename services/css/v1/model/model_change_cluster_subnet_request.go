package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeClusterSubnetRequest Request Object
type ChangeClusterSubnetRequest struct {

	// **参数解释**： 待切换子网的集群ID。获取方法请参见[获取集群ID](css_03_0101.xml)。 **约束限制**： 不涉及 **取值范围**： 集群ID。 **默认取值**： 不涉及
	ClusterId string `json:"cluster_id"`

	Body *ClusterChangeMainSubnet `json:"body,omitempty"`
}

func (o ChangeClusterSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeClusterSubnetRequest struct{}"
	}

	return strings.Join([]string{"ChangeClusterSubnetRequest", string(data)}, " ")
}
