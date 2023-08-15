package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClusterReq struct {

	// 新的集群名称。
	ClusterName string `json:"cluster_name"`
}

func (o UpdateClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterReq struct{}"
	}

	return strings.Join([]string{"UpdateClusterReq", string(data)}, " ")
}
