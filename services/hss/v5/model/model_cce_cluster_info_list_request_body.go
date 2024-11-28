package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceClusterInfoListRequestBody CCE集群配置请求
type CceClusterInfoListRequestBody struct {

	// 集群id列表
	ClusterInfoList []CceClusterInfoListRequestBodyClusterInfoList `json:"cluster_info_list"`

	// 集群id列表
	ClusterIdList *[]string `json:"cluster_id_list,omitempty"`
}

func (o CceClusterInfoListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceClusterInfoListRequestBody struct{}"
	}

	return strings.Join([]string{"CceClusterInfoListRequestBody", string(data)}, " ")
}
