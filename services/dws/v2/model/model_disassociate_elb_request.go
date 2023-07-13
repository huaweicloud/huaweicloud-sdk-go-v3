package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateElbRequest Request Object
type DisassociateElbRequest struct {

	// 集群ID
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
