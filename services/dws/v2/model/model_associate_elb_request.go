package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateElbRequest Request Object
type AssociateElbRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 未绑定的弹性负载均衡ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ElbId string `json:"elb_id"`
}

func (o AssociateElbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateElbRequest struct{}"
	}

	return strings.Join([]string{"AssociateElbRequest", string(data)}, " ")
}
