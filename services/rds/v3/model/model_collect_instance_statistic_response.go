package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectInstanceStatisticResponse Response Object
type CollectInstanceStatisticResponse struct {

	// 实例总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 异常实例数
	AbnormalNum *int32 `json:"abnormal_num,omitempty"`

	// 磁盘不足实例数
	DiskFullNum *int32 `json:"disk_full_num,omitempty"`

	// 冻结实例数
	FrozenNum *int32 `json:"frozen_num,omitempty"`

	// 运行中实例数
	NormalNum *int32 `json:"normal_num,omitempty"`

	// 等待重启实例数
	WaitRebootNum  *int32 `json:"wait_reboot_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CollectInstanceStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectInstanceStatisticResponse struct{}"
	}

	return strings.Join([]string{"CollectInstanceStatisticResponse", string(data)}, " ")
}
