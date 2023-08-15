package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpcepWhitelistRequest Request Object
type UpdateVpcepWhitelistRequest struct {

	// 指定待更改的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *UpdateVpcepWhitelistReq `json:"body,omitempty"`
}

func (o UpdateVpcepWhitelistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpcepWhitelistRequest struct{}"
	}

	return strings.Join([]string{"UpdateVpcepWhitelistRequest", string(data)}, " ")
}
