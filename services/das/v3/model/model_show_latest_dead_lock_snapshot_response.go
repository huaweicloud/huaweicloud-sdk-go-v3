package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestDeadLockSnapshotResponse Response Object
type ShowLatestDeadLockSnapshotResponse struct {
	MysqlDeadLock *QueryLatestDeadLockRespMysqlDeadLock `json:"mysql_dead_lock,omitempty"`

	// 表示系统是否成功检测到死锁事件
	FindLock *bool `json:"find_lock,omitempty"`

	// 死锁事件发生的 Unix 时间戳
	HappenTime     *int32 `json:"happen_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowLatestDeadLockSnapshotResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestDeadLockSnapshotResponse struct{}"
	}

	return strings.Join([]string{"ShowLatestDeadLockSnapshotResponse", string(data)}, " ")
}
