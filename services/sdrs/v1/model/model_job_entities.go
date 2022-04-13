package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type JobEntities struct {
	// 复制对ID

	ReplicationPairId *string `json:"replication_pair_id,omitempty"`
	// 组成复制对的云硬盘ID

	VolumeIds *string `json:"volume_ids,omitempty"`
	// 保护组ID

	ServerGroupId *string `json:"server_group_id,omitempty"`
	// 保护实例ID

	ProtectedInstanceId *string `json:"protected_instance_id,omitempty"`
	// 容灾站点服务器ID

	NativeServerId *string `json:"native_server_id,omitempty"`
	// 网卡ID

	NicId *string `json:"nic_id,omitempty"`
	// 子Job信息

	SubJobs *[]SubJobParams `json:"sub_jobs,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
