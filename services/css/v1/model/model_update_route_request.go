package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRouteRequest Request Object
type UpdateRouteRequest struct {

	// 指定待操作的集群ID。
	ClusterId string `json:"cluster_id"`

	Body *UpdateRouteRequestBody `json:"body,omitempty"`
}

func (o UpdateRouteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRouteRequest struct{}"
	}

	return strings.Join([]string{"UpdateRouteRequest", string(data)}, " ")
}
