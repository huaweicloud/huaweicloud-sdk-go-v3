package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2CreateClusterReq V2接口创建集群请求体。
type V2CreateClusterReq struct {
	Cluster *V2CreateCluster `json:"cluster,omitempty"`
}

func (o V2CreateClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2CreateClusterReq struct{}"
	}

	return strings.Join([]string{"V2CreateClusterReq", string(data)}, " ")
}
