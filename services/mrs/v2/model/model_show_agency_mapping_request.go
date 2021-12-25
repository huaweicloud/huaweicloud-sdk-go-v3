package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowAgencyMappingRequest struct {
	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。

	ClusterId string `json:"cluster_id"`
}

func (o ShowAgencyMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyMappingRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyMappingRequest", string(data)}, " ")
}
