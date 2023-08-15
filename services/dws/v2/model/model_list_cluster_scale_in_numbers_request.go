package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterScaleInNumbersRequest Request Object
type ListClusterScaleInNumbersRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterScaleInNumbersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterScaleInNumbersRequest struct{}"
	}

	return strings.Join([]string{"ListClusterScaleInNumbersRequest", string(data)}, " ")
}
