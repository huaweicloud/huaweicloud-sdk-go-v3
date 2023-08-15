package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateElbRequest Request Object
type AssociateElbRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 未绑定的弹性负载均衡ID
	ElbId string `json:"elb_id"`
}

func (o AssociateElbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateElbRequest struct{}"
	}

	return strings.Join([]string{"AssociateElbRequest", string(data)}, " ")
}
