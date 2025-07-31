package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceClusterIdListRequestBody CCE集群配置请求
type CceClusterIdListRequestBody struct {

	// 集群id列表
	ClusterIdList *[]string `json:"cluster_id_list,omitempty"`
}

func (o CceClusterIdListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceClusterIdListRequestBody struct{}"
	}

	return strings.Join([]string{"CceClusterIdListRequestBody", string(data)}, " ")
}
