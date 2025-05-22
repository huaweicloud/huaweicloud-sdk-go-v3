package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateElbRequest Request Object
type DisassociateElbRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// 集群已绑定的弹性负载均衡ID
	ElbId string `json:"elb_id"`
}

func (o DisassociateElbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateElbRequest struct{}"
	}

	return strings.Join([]string{"DisassociateElbRequest", string(data)}, " ")
}
