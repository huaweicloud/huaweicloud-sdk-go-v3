package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CceClusterInfoListRequestBodyClusterInfoList struct {

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName string `json:"cluster_name"`
}

func (o CceClusterInfoListRequestBodyClusterInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceClusterInfoListRequestBodyClusterInfoList struct{}"
	}

	return strings.Join([]string{"CceClusterInfoListRequestBodyClusterInfoList", string(data)}, " ")
}
