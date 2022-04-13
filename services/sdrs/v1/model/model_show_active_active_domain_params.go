package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询双活域响应体结构
type ShowActiveActiveDomainParams struct {
	// 双活域ID。

	Id string `json:"id"`
	// 双活域名称。

	Name string `json:"name"`
	// 双活域描述。

	Description string `json:"description"`
	// 表示该双活域下的资源是否售罄。

	SoldOut bool `json:"sold_out"`

	LocalReplicationCluster *ReplicationClusterParams `json:"local_replication_cluster"`

	RemoteReplicationCluster *ReplicationClusterParams `json:"remote_replication_cluster"`
}

func (o ShowActiveActiveDomainParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActiveActiveDomainParams struct{}"
	}

	return strings.Join([]string{"ShowActiveActiveDomainParams", string(data)}, " ")
}
