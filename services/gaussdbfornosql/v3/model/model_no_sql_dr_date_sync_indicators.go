package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 双活实例rsync同步状态指标值
type NoSqlDrDateSyncIndicators struct {

	// 节点内同步命令的执行速率,每秒多少条数据；
	RsyncOps *int64 `json:"rsync_ops,omitempty"`

	// 节点内的同步WAL堆积大小,单位MB；
	RsyncWalSize *int64 `json:"rsync_wal_size,omitempty"`

	// 同步消息从放入消息队列，直到收到对端响应的平均耗时，单位us；
	RsyncPushCost *int64 `json:"rsync_push_cost,omitempty"`

	// 同步消息从消息队列取出，直到收到对端响应的平均耗时，单位us；
	RsyncSendCost *int64 `json:"rsync_send_cost,omitempty"`

	// 采集周期内rsync的同步推送耗时最大值，单位us;
	RsyncMaxPushCost *int64 `json:"rsync_max_push_cost,omitempty"`

	// 采集周期内rsync的同步发送耗时最大值，单位us;
	RsyncMaxSendCost *int64 `json:"rsync_max_send_cost,omitempty"`

	// rsync的同步状态，1表示正在同步，0表示没有同步;
	RsyncStatus *int32 `json:"rsync_status,omitempty"`
}

func (o NoSqlDrDateSyncIndicators) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoSqlDrDateSyncIndicators struct{}"
	}

	return strings.Join([]string{"NoSqlDrDateSyncIndicators", string(data)}, " ")
}
