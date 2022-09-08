package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteEdgeCloudRequest struct {

	// 边缘业务ID。
	EdgecloudId string `json:"edgecloud_id"`
}

func (o DeleteEdgeCloudRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeCloudRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeCloudRequest", string(data)}, " ")
}
