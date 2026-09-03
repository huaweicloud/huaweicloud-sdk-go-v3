package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestDeadLockSnapshot4ApiResponse Response Object
type ShowLatestDeadLockSnapshot4ApiResponse struct {

	// 是否找到有锁
	FindLock *bool `json:"find_lock,omitempty"`

	// 发生时间（ms）
	HappenTime *int64 `json:"happen_time,omitempty"`

	MysqlDeadLock  *MySqlDeadLock `json:"mysql_dead_lock,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowLatestDeadLockSnapshot4ApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestDeadLockSnapshot4ApiResponse struct{}"
	}

	return strings.Join([]string{"ShowLatestDeadLockSnapshot4ApiResponse", string(data)}, " ")
}
