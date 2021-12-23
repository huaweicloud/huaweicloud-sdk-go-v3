package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowEdgeCloudRequest struct {
	// 边缘业务ID。

	EdgecloudId string `json:"edgecloud_id"`
}

func (o ShowEdgeCloudRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeCloudRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeCloudRequest", string(data)}, " ")
}
