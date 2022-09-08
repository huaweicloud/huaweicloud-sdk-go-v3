package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateLinkRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 连接名称
	LinkName string `json:"link_name"`

	Body *CdmCreateAndUpdateLinkReq `json:"body,omitempty"`
}

func (o UpdateLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLinkRequest struct{}"
	}

	return strings.Join([]string{"UpdateLinkRequest", string(data)}, " ")
}
