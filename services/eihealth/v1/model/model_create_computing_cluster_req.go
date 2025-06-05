package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateComputingClusterReq 绑定计算集群请求体。
type CreateComputingClusterReq struct {

	// CCE集群ID。
	CceClusterId string `json:"cce_cluster_id"`
}

func (o CreateComputingClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComputingClusterReq struct{}"
	}

	return strings.Join([]string{"CreateComputingClusterReq", string(data)}, " ")
}
