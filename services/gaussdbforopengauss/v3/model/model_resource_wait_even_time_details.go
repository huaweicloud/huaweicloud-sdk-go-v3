package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceWaitEvenTimeDetails 资源类等待事件耗时
type ResourceWaitEvenTimeDetails struct {
	DataIoTime *DataIoTime `json:"data_io_time"`

	LockTime *LockTime `json:"lock_time"`

	LwlockTime *LwlockTime `json:"lwlock_time"`
}

func (o ResourceWaitEvenTimeDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceWaitEvenTimeDetails struct{}"
	}

	return strings.Join([]string{"ResourceWaitEvenTimeDetails", string(data)}, " ")
}
