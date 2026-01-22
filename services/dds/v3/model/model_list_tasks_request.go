package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksRequest Request Object
type ListTasksRequest struct {

	// 查询开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始，Z指时区偏移量
	StartTime string `json:"start_time"`

	// 查询结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于查询开始时间，时间跨度不超过30天。 其中，T指某个时间的开始，Z指时区偏移量。
	EndTime string `json:"end_time"`

	// 任务状态： 取值为“Running”为执行中； 取值为“Completed”为完成； 取值为“Failed” 为失败。
	Status *string `json:"status,omitempty"`

	// 任务ID。
	Id *string `json:"id,omitempty"`

	// 任务名称。对应取值如下： - \"CreateMongoDB\"：创建集群实例 - \"CreateMongoDBReplica\"：创建副本集实例 - \"CreateMongoDBReplicaSingle\"：创建单节点实例 - \"EnlargeMongoDBVolume\"：磁盘扩容 - \"ResizeMongoDBInstance\"：社区版实例规格变更 - \"ResizeDfvMongoDBInstance\"：社区增强版实例规格变更 - \"EnlargeMongoDBGroup\"：添加节点 - \"ReplicaSetEnlargeNode\"：副本集添加备节点 - \"AddReadonlyNode\"：添加只读节点 - \"RestartInstance\"：重启集群实例 - \"RestartGroup\"：重启集群节点组 - \"RestartNode\"：重启集群节点 - \"RestartReplicaSetInstance\"：重启副本集实例 - \"RestartReplicaSingleInstance\"：重启单节点实例 - \"SwitchPrimary\"：主备切换 - \"ModifyIp\"：修改内网地址 - \"ModifySecurityGroup\"：修改安全组 - \"ModifyPort\"：修改数据库端口 - \"BindPublicIP\"：绑定弹性IP - \"UnbindPublicIP\"：解绑弹性IP - \"SwitchInstanceSSL\"：切换SSL - \"AzMigrate\"：迁移可用区 - \"CreateIp\"：显示shard/config IP - \"ModifyOpLogSize\"：修改oplog大小 - \"RestoreMongoDB\"：集群恢复到新实例 - \"RestoreMongoDB_Replica\"：副本集恢复到新实例 - \"RestoreMongoDB_Replica_Single\"：单节点恢复到新实例 - \"RestoreMongoDB_Replica_PITR\"：副本集恢复到指定时间点 - \"MongodbSnapshotBackup\"：创建物理备份 - \"MongodbSnapshotEBackup\"：创建快照备份 - \"MongodbRestoreData2CurrentInstance\"：备份恢复到当前实例 - \"MongodbRestoreData2NewInstance\"：备份恢复到新实例 - \"MongodbPitr2CurrentInstance\"：备份恢复到当前实例指定时间点 - \"MongodbPitr2NewInstance\"：备份恢复到新实例指定时间点 - \"MongodbRecycleBackup\"：备份回收 - \"MongodbRestoreTable\"：库表级时间点恢复 - \"UpgradeDatabaseVersion\"：升级数据库补丁
	Name *string `json:"name,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}
