package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowLinkRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 连接名称
	LinkName string `json:"link_name" xml:"link_name"`
}

func (o ShowLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowLinkRequest", string(data)}, " ")
}
