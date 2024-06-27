package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaptureTaskVo struct {

	// 抓包大小
	CaptureSize *string `json:"capture_size,omitempty"`

	// 抓包创建时间
	CreatedDate *string `json:"created_date,omitempty"`

	// 目的地址
	DestAddress *string `json:"dest_address,omitempty"`

	// 目的地址类型0 ipv4,1 ipv6
	DestAddressType *int32 `json:"dest_address_type,omitempty"`

	// 目的端口
	DestPort *string `json:"dest_port,omitempty"`

	// 抓包时长
	Duration *int32 `json:"duration,omitempty"`

	// 是否被删除，0否 1是
	IsDeleted *int32 `json:"is_deleted,omitempty"`

	// 最大抓包数
	MaxPackets *int32 `json:"max_packets,omitempty"`

	// 修改日期
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 抓包任务名称
	Name *string `json:"name,omitempty"`

	// 协议类型:TCP为6，UDP为17，ICMP为1，ICMPV6为58，ANY为-1，手动类型不为空，自动类型为空
	Protocol *int32 `json:"protocol,omitempty"`

	// 剩余保留天数
	RemainingDays *int32 `json:"remaining_days,omitempty"`

	// 源地址
	SourceAddress *string `json:"source_address,omitempty"`

	// 源地址类型0 ipv4,1 ipv6
	SourceAddressType *int32 `json:"source_address_type,omitempty"`

	// 源端口
	SourcePort *string `json:"source_port,omitempty"`

	// 抓包任务状态
	Status *int32 `json:"status,omitempty"`

	// 抓包任务id
	TaskId *string `json:"task_id,omitempty"`
}

func (o CaptureTaskVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaptureTaskVo struct{}"
	}

	return strings.Join([]string{"CaptureTaskVo", string(data)}, " ")
}
