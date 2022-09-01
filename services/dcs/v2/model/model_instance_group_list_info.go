package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分片列表
type InstanceGroupListInfo struct {

	// 分片id
	GroupId *string `json:"group_id,omitempty" xml:"group_id"`

	// 分片名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 每个分片包含的副本列表。
	ReplicationList *[]InstanceReplicationListInfo `json:"replication_list,omitempty" xml:"replication_list"`
}

func (o InstanceGroupListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceGroupListInfo struct{}"
	}

	return strings.Join([]string{"InstanceGroupListInfo", string(data)}, " ")
}
