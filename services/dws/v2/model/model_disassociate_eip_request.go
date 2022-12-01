package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisassociateEipRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 集群绑定的弹性IP
	EipId string `json:"eip_id"`
}

func (o DisassociateEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateEipRequest struct{}"
	}

	return strings.Join([]string{"DisassociateEipRequest", string(data)}, " ")
}
