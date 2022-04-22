package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateClusterNameRequest struct {

	// 指定待更改的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *UpdateClusterNameReq `json:"body,omitempty"`
}

func (o UpdateClusterNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterNameRequest", string(data)}, " ")
}
