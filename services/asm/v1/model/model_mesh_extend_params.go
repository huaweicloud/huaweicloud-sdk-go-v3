package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MeshExtendParams struct {

	// 网格内集群信息
	Clusters []MeshCluster `json:"clusters"`
}

func (o MeshExtendParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshExtendParams struct{}"
	}

	return strings.Join([]string{"MeshExtendParams", string(data)}, " ")
}
