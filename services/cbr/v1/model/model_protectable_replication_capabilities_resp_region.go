package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProtectableReplicationCapabilitiesRespRegion struct {

	// 云服务所在的区域
	Name string `json:"name"`

	// 支持复制的目标区域列表
	ReplicationDestinations []string `json:"replication_destinations"`
}

func (o ProtectableReplicationCapabilitiesRespRegion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectableReplicationCapabilitiesRespRegion struct{}"
	}

	return strings.Join([]string{"ProtectableReplicationCapabilitiesRespRegion", string(data)}, " ")
}
