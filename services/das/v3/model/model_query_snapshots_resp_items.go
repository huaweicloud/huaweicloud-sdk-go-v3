package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuerySnapshotsRespItems struct {

	// snapshot编号
	Id *int32 `json:"id,omitempty"`

	// 表示状态，0表示等待，1表示正在运行，2表示失败，3表示成功
	Status *int32 `json:"status,omitempty"`

	// 创建时间
	CreateAt *int32 `json:"create_at,omitempty"`

	// 是否发现锁，1表示发现了锁，0表示没有
	FindLock *int32 `json:"find_lock,omitempty"`

	// 发生锁的时间
	HappenAt *int32 `json:"happen_at,omitempty"`
}

func (o QuerySnapshotsRespItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySnapshotsRespItems struct{}"
	}

	return strings.Join([]string{"QuerySnapshotsRespItems", string(data)}, " ")
}
