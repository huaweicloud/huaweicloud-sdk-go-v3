package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceDrInfo struct {

	// 容灾关系id
	Id string `json:"id"`

	// 容灾搭建状态
	Status string `json:"status"`

	// 失败消息
	FailedMessage *string `json:"failed_message,omitempty"`

	// 同步状态，取值范围是0或-1，0表示正常，-1表示异常。
	ReplicaState *string `json:"replica_state,omitempty"`

	// 发送延迟大小（MB），即主实例当前wal日志写入位点与灾备实例当前接收wal日志位点的差值。
	WalWriteReceiveDelayInMb *string `json:"wal_write_receive_delay_in_mb,omitempty"`

	// 端到端延迟大小（MB），即主实例当前wal日志写入位点与灾备实例当前回放wal日志位点的差值。
	WalWriteReplayDelayInMb *string `json:"wal_write_replay_delay_in_mb,omitempty"`

	// 回放延迟时间（ms），即数据在灾备上回放的延迟时间。
	WalReceiveReplayDelayInMs *string `json:"wal_receive_replay_delay_in_ms,omitempty"`

	// 主实例ID
	MasterInstanceId string `json:"master_instance_id"`

	// 主实例所在region
	MasterRegion string `json:"master_region"`

	// 灾备实例ID
	SlaveInstanceId string `json:"slave_instance_id"`

	// 灾备实例所在region
	SlaveRegion string `json:"slave_region"`

	// 灾备搭建时间
	Time int64 `json:"time"`
}

func (o InstanceDrInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDrInfo struct{}"
	}

	return strings.Join([]string{"InstanceDrInfo", string(data)}, " ")
}
