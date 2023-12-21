package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReplicationInfo struct {

	// 副本ID
	ReplicationId *string `json:"replication_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 副本IP
	ReplicationIp *string `json:"replication_ip,omitempty"`

	// 组ID
	GroupId *string `json:"group_id,omitempty"`

	// 组名称
	GroupName *string `json:"group_name,omitempty"`

	// 可用区
	AvailableZone *string `json:"available_zone,omitempty"`
}

func (o ReplicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplicationInfo struct{}"
	}

	return strings.Join([]string{"ReplicationInfo", string(data)}, " ")
}
