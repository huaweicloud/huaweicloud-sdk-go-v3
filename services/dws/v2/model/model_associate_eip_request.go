package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateEipRequest Request Object
type AssociateEipRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 未绑定的弹性IP的ID
	EipId string `json:"eip_id"`
}

func (o AssociateEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateEipRequest struct{}"
	}

	return strings.Join([]string{"AssociateEipRequest", string(data)}, " ")
}
