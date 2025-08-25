package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCcspClusterRequest Request Object
type ShowCcspClusterRequest struct {

	// 密码服务集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ShowCcspClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCcspClusterRequest struct{}"
	}

	return strings.Join([]string{"ShowCcspClusterRequest", string(data)}, " ")
}
