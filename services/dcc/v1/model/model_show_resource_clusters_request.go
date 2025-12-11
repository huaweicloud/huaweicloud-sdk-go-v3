package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceClustersRequest Request Object
type ShowResourceClustersRequest struct {

	// 集群服务类型，取值范围：“ecs”或“bms”。
	ServiceType *string `json:"service_type,omitempty"`
}

func (o ShowResourceClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceClustersRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceClustersRequest", string(data)}, " ")
}
