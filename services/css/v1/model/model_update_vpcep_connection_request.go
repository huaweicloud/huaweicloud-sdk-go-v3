package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateVpcepConnectionRequest struct {
	// 指定待更新的集群ID。

	ClusterId string `json:"cluster_id"`

	Body *UpdateVpcepConnectionReq `json:"body,omitempty"`
}

func (o UpdateVpcepConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcepConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpcepConnectionRequest", string(data)}, " ")
}
