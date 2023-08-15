package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppInstancesRequest Request Object
type ListAppInstancesRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListAppInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAppInstancesRequest", string(data)}, " ")
}
