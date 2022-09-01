package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowClusterDetailsRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`
}

func (o ShowClusterDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterDetailsRequest", string(data)}, " ")
}
