package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCcspClusterRequest Request Object
type DeleteCcspClusterRequest struct {

	// 密码服务集群ID
	ClusterId string `json:"cluster_id"`
}

func (o DeleteCcspClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCcspClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteCcspClusterRequest", string(data)}, " ")
}
